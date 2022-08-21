a: 1
b: 2
_hidden: 3
sum: a + b
l: [a, b]

map: [string]:int
map: {a: 1*5}
map: {"b": b*5}

// schema
#Person: {
	name:  string
	age:   int
	human: true
	// optional field
	enrolled?: bool
}

ben: #Person & {
	name: "Ben"
	age:  31
}
