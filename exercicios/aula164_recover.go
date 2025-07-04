package main

import "fmt"

func proteger() {
	if r := recover(); r != nil {
		fmt.Println("Recuperado de um panic:", r)
	}
}

func podeFalhar() {
	defer proteger()
	panic("algo deu muito errado!")
}

func main() {
	podeFalhar()
	fmt.Println("Programa continua executando normalmente.")
} 