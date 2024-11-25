# cue_config Data Source

`cue_config` allows Terraform to evaluate [CUE](https://cuelang.org/docs/) configs and render JSON for usage.

## Usage

Define a `cue_config` data source to validate CUE `content`.

```tf
data "cue_config" "example" {
  content = <<EOF
a: 1
b: 2
sum: a + b
_hidden: 3
l: [a, b]

map: [string]:int
map: {a: 1 * 5}
map: {"b": b * 5}
EOF
  pretty_print = true
}
```

Alternately, provide `paths` to CUE files (supports imports).

```tf
data "cue_config" "example" {
  paths = [
    "core.cue",
    "box.cue",
  ]
  # `paths` and `content` can be combined
  # content = "foo: bar"
  pretty_print = false
}
```

Customize the root directory CUE uses during loading (defaults to current directory).

```
data "cue_config" "example" {
  paths = [
    "foo.cue",
  ]
  dir = "internal/testmod"
  pretty_print = false
}
```

Render the CUE config as JSON for use in Terraform expressions.

```tf
output "out" {
  description = "Show Cue rendered as JSON"
  value       = data.cue_config.example.rendered
}
```

## Argument Reference

* `content` - inline CUE contents to evaluate
* `dir` - Root directory CUE uses to load and evaluate imports
* `paths` - list of paths to CUE files (relative to Terraform workspace) to evaluate (exclusive with `content`)
* `expression` - evaluate an expression instead of an entire config
* `pretty_print` - indent rendered JSON for visual prettiness (default: false)

## Argument Attributes

* `rendered` - JSON output from evaluating CUE content(s)

