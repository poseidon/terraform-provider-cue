provider "cue" {}

terraform {
  required_providers {
    cue = {
      source  = "poseidon/cue"
      version = "0.2.0"
    }
  }
}
