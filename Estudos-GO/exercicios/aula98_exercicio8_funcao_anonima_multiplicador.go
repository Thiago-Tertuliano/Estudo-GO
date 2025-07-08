package main

import "fmt"

func main() {
	numeros := []int{7, 7}

	multiplicador := func(numero []int) int {
		total := 1
		for _, valor := range numero {
			total *= valor
		}
		return total
	}

	resultado := multiplicador(numeros)
	fmt.Println(resultado)
} 