package main

import "fmt"

func executarOperacao(a, b int, operacao func(int, int) int) {
	resultado := operacao(a, b)
	fmt.Println("Resultado:", resultado)
}

func main() {
	soma := func(x, y int) int {
		return x + y
	}

	multiplicacao := func(x, y int) int {
		return x * y
	}

	executarOperacao(3, 4, soma)
	executarOperacao(3, 4, multiplicacao)
} 