# terraform-provider-cue

Notable changes between releases.

## Latest

## v0.2.0

* Add `expression` to `cue_config` ([#3](https://github.com/poseidon/terraform-provider-cue/pull/3))
  * Evaluate an expression instead of an entire config
* Add `dir` to `cue_config` to set the loader
  * Pass a root dir to load and evaluate imports

## v0.1.0

* Add `cue_config` data source to evaluate CUE contents
  * Evaluate inline CUE `content`
  * Evaluate CUE files specified via `paths` list
  * Support `pretty_print` option (default false)
  * Render CUE values as JSON
