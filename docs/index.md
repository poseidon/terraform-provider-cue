# terraform-provider-cue

`terraform-provider-cue` allows Terraform to evaluate [CUE](https://cuelang.org/docs/) configs and render JSON for use in Terraform.

## Usage

Configure the `cue` provider (e.g. `providers.tf`).

```tf
provider "cue" {}

terraform {
  required_providers {
    ct = {
      source  = "poseidon/cue"
      version = "0.4.0"
    }
  }
}
```

Run `terraform init` to ensure version requirements are met.

```
$ terraform init -upgrade
```

Define a `cue_config` data source to validate CUE `content`.

```tf
data "cue_config" "example" {
  pretty_print = true
  content = <<-EOT
    a: 1
    b: 2
    sum: a + b
    _hidden: 3
    l: [a, b]

    map: [string]:int
    map: {a: 1 * 5}
    map: {"b": b * 5}
  EOT
}
```

Optionally provide `paths` to CUE files (supports imports).

```tf
data "cue_config" "example" {
  paths = [
    "core.cue",
    "box.cue",
  ]
}
```

Or unify `content` and `path` based expressions together.

```tf
data "cue_config" "example" {
  paths = [
    "partial.cue",
  ]
  content = <<-EOT
    package example

    _config: {
      name: "ACME"
      amount: "$20.00"
    }
  EOT
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

The rendered `content` example looks like:

```json
{
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
}
```

## Requirements

* Terraform v1.0+ [installed](https://www.terraform.io/downloads.html)

## Development

### Binary

To develop the provider plugin locally, build an executable with Go v1.18+.

```
make
```

