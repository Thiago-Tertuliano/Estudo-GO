package main

import "fmt"

var x [5]int
var y [6]int

func main() {
	x[0] = 1
	x[1] = 10
	x[2] = 20
	x[3] = 30
	x[4] = 40
	fmt.Println(x[0], x[1], x[2], x[3], x[4])
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(len(x))
} 