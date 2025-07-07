package main

import "fmt"

func main() {
	x := 10
	y := 20

	if x < y && y > 15 {
		fmt.Println("x é menor que y e y é maior que 15")
	}

	if x > 5 || y < 25 {
		fmt.Println("x é maior que 5 ou y é menor que 25")
	}

	if !(x == y) {
		fmt.Println("x não é igual a y")
	}
} 