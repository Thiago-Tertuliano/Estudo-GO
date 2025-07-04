package main

import "fmt"

func main() {
	var x int = 10
	var y int = 20

	fmt.Printf("%d == %d", x, y)
	s := x == y
	fmt.Println("\n", s)
	fmt.Printf("\n%d >= %d", x, y)
	a := x >= y
	fmt.Println("\n", a)
	fmt.Printf("\n%d <= %d", x, y)
	b := x <= y
	fmt.Println("\n", b)
	fmt.Printf("\n%d > %d", x, y)
	c := x > y
	fmt.Println("\n", c)
	fmt.Printf("\n%d < %d", x, y)
	d := x < y
	fmt.Println("\n", d)
} 