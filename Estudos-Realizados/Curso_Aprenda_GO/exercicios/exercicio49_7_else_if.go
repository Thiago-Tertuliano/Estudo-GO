package main

import "fmt"

func main() {
	var idade int = 21

	if idade < 18 {
		fmt.Printf("O jovem não pode ser escalado!")
	} else if idade > 20 {
		fmt.Printf("O jovem pode ser escalado e jogar")
	} else {
		fmt.Printf("O jovem pode jogar, mas não será escalado")
	}
} 