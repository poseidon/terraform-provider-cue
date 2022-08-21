package schema

#Box: {
  color: string
  row: int | *0
  column: int | *0
}

#Layout: {
  #boxes: [...#Box]
  boxes: [for i, box in #boxes {
    box
    row: div(i, 2)
    column: rem(i, 2)
  }]
}
