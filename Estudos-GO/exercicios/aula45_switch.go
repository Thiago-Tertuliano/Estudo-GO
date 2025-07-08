package main

import "fmt"

func main() {
	x := 10

	switch x {
	case 5:
		fmt.Println("x é igual a 5")
	case 10:
		fmt.Println("x é igual a 10")
	case 15:
		fmt.Println("x é igual a 15")
	default:
		fmt.Println("x não é igual a nenhum dos casos anteriores")
	}

	// Exemplo de switch com múltiplos casos
	switch x {
	case 5, 10, 15:
		fmt.Println("x é 5, 10 ou 15")
	default:
		fmt.Println("x não é 5, 10 ou 15")
	}
} 