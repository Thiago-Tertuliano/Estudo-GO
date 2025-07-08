package main

import "fmt"

func main() {
	// Loop for com inicialização, condição e pós-ação
	for i := 0; i < 5; i++ {
		fmt.Println("Iteração:", i)
	}

	// Loop while (usando for sem inicialização e pós-ação)
	j := 0
	for j < 5 {
		fmt.Println("Iteração while:", j)
		j++
	}

	// Loop infinito (comentado para evitar execução infinita)
	// k := 0
	// for {
	// 	fmt.Println("Loop infinito:", k)
	// 	k++
	// 	if k >= 5 {
	// 		break
	// 	}
	// }
} 