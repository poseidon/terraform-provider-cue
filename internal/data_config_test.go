package internal

import (
	"regexp"
	"testing"

	r "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const cueWithContent = `
data "cue_config" "example" {
	content = <<EOT
a: 1
b: 2
_hidden: 3
sum: a + b
l: [a, b]

map: [string]:int
map: {a: 1*5}
map: {"b": b*5}
EOT
}
`

const cueWithContentPretty = `
data "cue_config" "example" {
	content = <<EOT
a: 1
b: 2
_hidden: 3
sum: a + b
l: [a, b]

map: [string]:int
map: {a: 1*5}
map: {"b": b*5}
EOT
	pretty_print = true
}
`

const outputWithContent = `{"a":1,"b":2,"sum":3,"l":[1,2],"map":{"a":5,"b":10}}`
const outputWithContentPretty = `{
  "a": 1,
  "b": 2,
  "sum": 3,
  "l": [
    1,
    2
  ],
  "map": {
    "a": 5,
    "b": 10
  }
}`

const cueWithPaths = `
data "cue_config" "example" {
	paths = [
		"../examples/core.cue",
		# uses imports
		"../examples/box.cue",
	]
}
`

const outputWithPaths = `{"a":1,"b":2,"sum":3,"l":[1,2],"layout":{"boxes":[{"color":"red","row":0,"column":0},{"color":"blue","row":0,"column":1},{"color":"green","row":1,"column":0},{"color":"yellow","row":1,"column":1}]},"map":{"a":5,"b":10},"ben":{"name":"Ben","age":31,"human":true}}`

func TestConfigRender(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		Providers: testProviders,
		Steps: []r.TestStep{
			{
				Config: cueWithContent,
				Check: r.ComposeTestCheckFunc(
					r.TestCheckResourceAttr("data.cue_config.example", "rendered", outputWithContent),
				),
			},
			{
				Config: cueWithContentPretty,
				Check: r.ComposeTestCheckFunc(
					r.TestCheckResourceAttr("data.cue_config.example", "rendered", outputWithContentPretty),
				),
			},
			{
				Config: cueWithPaths,
				Check: r.ComposeTestCheckFunc(
					r.TestCheckResourceAttr("data.cue_config.example", "rendered", outputWithPaths),
				),
			},
		},
	})
}

func TestConfigContentOrPaths(t *testing.T) {
	hcl := `
data "cue_config" "invalid" {
	content = "a: 1"
	paths = [
		"../examples/core.cue",
	]
}
`
	r.UnitTest(t, r.TestCase{
		Providers: testProviders,
		Steps: []r.TestStep{
			{
				Config:      hcl,
				ExpectError: regexp.MustCompile("are mutually exclusive"),
			},
		},
	})
}
