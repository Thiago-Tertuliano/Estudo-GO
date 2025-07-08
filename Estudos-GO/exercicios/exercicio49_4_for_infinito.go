package main

import "fmt"

func main() {
	nascimento := 2005
	anoAtual := 2025

	for {
		if nascimento > anoAtual {
			break
		}
		fmt.Println(nascimento)
		nascimento++
	}
} 