package main

import "fmt"

const (
	x = iota // x será 0
	y = iota // y será 1
	z = iota // z será 2
	_ = iota // _ será 3, mas não será usado
	thiago = iota // thiago será 4
)

func main() {
	fmt.Println(x, y, z, thiago) // Imprime: 0 1 2 4
} 