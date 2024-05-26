package internal

import (
	"testing"

	r "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const cueWithContent = `
data "cue_config" "example" {
	content = <<-EOT
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
	pretty_print = true
	content = <<-EOT
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

const outputWithPaths = `{"layout":{"boxes":[{"color":"red","row":0,"column":0},{"color":"blue","row":0,"column":1},{"color":"green","row":1,"column":0},{"color":"yellow","row":1,"column":1}]},"a":1,"b":2,"sum":3,"l":[1,2],"map":{"a":5,"b":10},"ben":{"name":"Ben","age":31,"human":true}}`

const cueWithContentAndPaths = `
data "cue_config" "example" {
	paths = [
		"../examples/partial.cue",
	]
	content = <<-EOT
		package examples

		_config: {
			name: "ACME"
			amount: "$20.00"
		}
	EOT
}
`
const outputWithContentAndPaths = `{"title":"Invoice","customer":"ACME","bill":"$20.00"}`

const cueWithDir = `
data "cue_config" "example" {
	paths = [
		"core.cue",
		# uses imports
		"box.cue",
	]
	dir = "../examples"
}
`

const cueWithDir2 = `
data "cue_config" "example" {
	paths = [
		"foo.cue",
	]
	dir = "testmod"
}
`

const cueWithExpression = `
data "cue_config" "example" {
	content = <<EOT
a: 1
b: 2
_hidden: 3
EOT
	expression = "a"
}
`

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
			{
				Config: cueWithContentAndPaths,
				Check: r.ComposeTestCheckFunc(
					r.TestCheckResourceAttr("data.cue_config.example", "rendered", outputWithContentAndPaths),
				),
			},
			{
				Config: cueWithDir,
				Check: r.ComposeTestCheckFunc(
					r.TestCheckResourceAttr("data.cue_config.example", "rendered", outputWithPaths),
				),
			},
			{
				Config: cueWithDir2,
				Check: r.ComposeTestCheckFunc(
					r.TestCheckResourceAttr("data.cue_config.example", "rendered", `{"a":1}`),
				),
			},
			{
				Config: cueWithExpression,
				Check: r.ComposeTestCheckFunc(
					r.TestCheckResourceAttr("data.cue_config.example", "rendered", "1"),
				),
			},
		},
	})
}
