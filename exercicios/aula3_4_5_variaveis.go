package main

import "fmt"

func main() {
	var nome string = "John"
	fmt.Println("Nome:", nome)

	idade := 30
	fmt.Println("Idade:", idade)

	cidade, pais := "São Caetano do Sul", "Brasil"
	fmt.Println("Cidade:", cidade, "\nPaís:", pais)

	const pi = 3.14
	fmt.Println("Pi:", pi)
} 