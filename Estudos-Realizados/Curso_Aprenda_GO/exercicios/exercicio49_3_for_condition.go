package main

import "fmt"

func main() {
	nascimento := 2005
	anoAtual := 2025

	for nascimento <= anoAtual {
		fmt.Println(nascimento)
		nascimento++
	}
} 