# terraform-provider-cue

Notable changes between releases.

## Latest

* Add `expression` field to `cue_config` ([#3](https://github.com/poseidon/terraform-provider-cue/pull/3))
  * Evaluate an expression instead of an entire config

## v0.1.0

* Add `cue_config` data source to evaluate CUE contents
  * Evaluate inline CUE `content`
  * Evaluate CUE files specified via `paths` list
  * Support `pretty_print` option (default false)
  * Render CUE values as JSON
