package main

import "fmt"

func main() {
	var idade int = 11

	if idade < 18 {
		fmt.Printf("O jovem não pode entrar!")
	} else {
		fmt.Printf("Pode entrar")
	}
} 