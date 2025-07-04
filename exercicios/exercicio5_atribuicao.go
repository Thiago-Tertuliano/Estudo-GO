package main

import "fmt"

type thiago int

var x thiago
var y thiago

func main() {
	fmt.Printf("%v, %T\n", x, x)
	x = 42
	fmt.Println(x)

	y = x
	fmt.Printf("%v, %T\n", y, y)
} 