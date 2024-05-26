# terraform-provider-cue

Notable changes between releases.

## Latest

## v0.4.0

* Allow unifying `content` and `path` based CUE expressions ([#48](https://github.com/poseidon/terraform-provider-cue/pull/48))

## v0.3.0

* Update CUE from v0.4.3 to [v0.5.0](https://github.com/cue-lang/cue/releases/tag/v0.5.0)

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
