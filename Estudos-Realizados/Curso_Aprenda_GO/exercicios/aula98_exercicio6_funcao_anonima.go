package main

import "fmt"

func main() {
	gol := func(acao string) {
		fmt.Printf("Olha o %s", acao)
	}

	gol("Golaço")
} 