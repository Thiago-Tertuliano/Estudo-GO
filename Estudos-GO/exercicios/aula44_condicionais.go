package main

import "fmt"

func main() {
	x := 10
	y := 20

	if x < y {
		fmt.Println("x é menor que y")
	} else if x > y {
		fmt.Println("x é maior que y")
	} else {
		fmt.Println("x é igual a y")
	}

	// Exemplo de uso de operadores lógicos
	if x < 15 && y > 15 {
		fmt.Println("x é menor que 15 e y é maior que 15")
	}

	if x == 10 || y == 20 {
		fmt.Println("x é igual a 10 ou y é igual a 20")
	}
} 