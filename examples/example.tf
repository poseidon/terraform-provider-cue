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
}

output "out" {
  description = "Show Cue rendered as JSON"
  value       = data.cue_config.example.rendered
}

data "cue_config" "example2" {
  paths = [
    "core.cue",
    "box.cue",
  ]
}


output "out2" {
  description = "Show Cue rendered as JSON"
  value       = data.cue_config.example2.rendered
}
