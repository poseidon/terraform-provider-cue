package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataConfig() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataConfigRead,
		Schema: map[string]*schema.Schema{
			"content": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"paths": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"dir": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"expression": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pretty_print": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"rendered": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "rendered JSON",
			},
		},
	}
}

func dataConfigRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// Cue has race conditions loading
	// https://github.com/cue-lang/cue/issues/460
	lock := meta.(*sync.Mutex)
	lock.Lock()
	defer lock.Unlock()

	var diags diag.Diagnostics
	pretty := d.Get("pretty_print").(bool)
	content := d.Get("content").(string)
	pathsAny := d.Get("paths").([]any)
	paths := make([]string, len(pathsAny))
	for i, path := range pathsAny {
		if path != nil {
			paths[i] = path.(string)
		}
	}

	// create a Cue context
	cuectx := cuecontext.New()

	var value cue.Value
	var err error
	if len(paths) < 1 {
		value = cuectx.CompileString(content)
	} else {
		// load cue "instances" from filesystem
		value, err = loadPaths(cuectx, d, content, paths)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// lookup an expression
	if expression, ok := d.Get("expression").(string); ok && expression != "" {
		value = value.LookupPath(cue.ParsePath(expression))
		if err := value.Err(); err != nil {
			return diag.FromErr(fmt.Errorf("expression error: %v", err))
		}
	}

	// Cue validate to surface better errors
	if err := value.Validate(); err != nil {
		return diag.FromErr(fmt.Errorf("validate error: %v", err))
	}

	rendered, err := marshalJSON(value, pretty)
	if err != nil {
		return diag.FromErr(fmt.Errorf("JSON marshal error: %v", err))
	}
	if err := d.Set("rendered", string(rendered)); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(hashcode(string(rendered)))
	return diags
}

// load Paths parses Cue files and merges them.
func loadPaths(cuectx *cue.Context, data *schema.ResourceData, content string, paths []string) (cue.Value, error) {
	dir := data.Get("dir").(string)

	config := &load.Config{
		Dir: dir,
		// Trick CUE into "loading" the content expression as though it was a file
		Overlay: map[string]load.Source{
			"/content.cue": load.FromString(content),
		},
	}
	// content.cue is a fake path to convince load to read string contents
	paths = append(paths, "/content.cue")

	// load cue "instances" from the given paths
	instances := load.Instances(paths, config)
	// Cue's API is clunky: We must inspect the instances to determine
	// if errors happened
	for _, inst := range instances {
		if inst.Err != nil {
			return cue.Value{}, inst.Err
		}
		if inst.Incomplete {
			return cue.Value{}, fmt.Errorf("incomplete load: %s", inst.PkgName)
		}
	}

	// build a Cue Value from each instance
	// Cue's API has poor naming and ergonomics
	values, err := cuectx.BuildInstances(instances)
	if err != nil {
		return cue.Value{}, err
	}

	// merge or "unify" Cue Values
	// Cue's API is buggy. Unifying with a zero value doesn't work.
	// https://github.com/cuelang/cue/issues/933#issuecomment-830760552
	/*
		var value cue.Value
		for _, val := range values {
			value.Unify(val)
		}
	*/
	value := values[0]
	for _, val := range values[1:] {
		value.Unify(val)
	}

	return value, nil
}

func marshalJSON(v any, pretty bool) ([]byte, error) {
	if pretty {
		return json.MarshalIndent(v, "", "  ")
	}
	return json.Marshal(v)
}
