package main

import "fmt"

var x int
var y int

func main() {
	x = 20
	fmt.Printf("Decimal: %d, Binário: %b, Hexadecimal: %#x\n", x, x, x)
	y = x << 1 // Deslocamento de bits para a esquerda
	fmt.Printf("Deslocamento para a esquerda: Decimal: %d, Binário: %b, Hexadecimal: %#x\n", y, y, y)
} 