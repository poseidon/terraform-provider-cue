provider "cue" {}

terraform {
  required_providers {
    cue = {
      source  = "poseidon/cue"
      #source = "terraform.localhost/poseidon/cue"
      version = "0.4.1"
    }
  }
}
