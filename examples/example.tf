data "cue_config" "example" {
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

data "cue_config" "example2" {
  paths = [
    "core.cue",
    "box.cue",
  ]
}

data "cue_config" "example3" {
  paths = [
    "partial.cue",
  ]
  content = <<-EOT
    package examples

		_config: {
			name: "ACME"
			amount: "$20.00"
		}
  EOT
}

output "out" {
  description = "Show Cue content rendered as JSON"
  value       = data.cue_config.example.rendered
}

output "out2" {
  description = "Show Cue files rendered as JSON"
  value       = data.cue_config.example2.rendered
}

output "out3" {
  description = "Show Cue content+files rendered as JSON"
  value       = data.cue_config.example3.rendered
}
