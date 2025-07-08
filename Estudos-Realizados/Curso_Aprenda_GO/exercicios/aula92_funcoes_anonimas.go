package main

import "fmt"

func main() {
	saudacao := func(nome string) {
		fmt.Printf("Olá, %s!\n", nome)
	}

	saudacao("Maria")

	func(a, b int) {
		fmt.Printf("A soma de %d e %d é %d\n", a, b, a+b)
	}(3, 4)
} 