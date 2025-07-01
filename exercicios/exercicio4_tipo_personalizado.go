package main

import "fmt"

type thiago int

var x thiago

func main() {
	fmt.Printf("%v, %T\n", x, x)
	x = 42
	fmt.Printf("%v", x)
} 