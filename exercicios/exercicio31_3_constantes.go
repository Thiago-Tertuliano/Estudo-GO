package main

import "fmt"

const x = 10 // Constante não tipada
const y int = 20 // Constante tipada

func main() {
	fmt.Println(x)
	fmt.Printf("%v, %T", y, y)
} 