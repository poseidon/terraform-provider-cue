import "github.com/poseidon/terraform-provider-cue/examples/schema"

layout: schema.#Layout & {
  #boxes: [
    schema.#Box & {color: "red"},
    schema.#Box & {color: "blue"},
    schema.#Box & {color: "green"},
    schema.#Box & {color: "yellow"},
  ]
}
