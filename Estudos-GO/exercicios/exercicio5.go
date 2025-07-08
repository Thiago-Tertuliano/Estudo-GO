package main

import "fmt"

type thiago int

var x thiago

func main() {
	x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
} 