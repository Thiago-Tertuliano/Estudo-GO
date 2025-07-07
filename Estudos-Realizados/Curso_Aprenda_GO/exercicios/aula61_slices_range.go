package main

import "fmt"

func main() {
	slice := []string{"bola", "chuteira", "luva"}

	for indice, valor := range slice {
		fmt.Println("No indice", indice, "tem a palavra: ", valor)
	}

	slice = append(slice, "melancia")
	for indice, valor := range slice {
		fmt.Println("No indice", indice, "tem a palavra: ", valor)
	}

	for _, valor := range slice {
		fmt.Println("Uma das palavras são: ", valor)
	}
} 