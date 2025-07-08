package main

import "fmt"

var a int
var b float64
var c string
var d bool

func main() {
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%q, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
} 