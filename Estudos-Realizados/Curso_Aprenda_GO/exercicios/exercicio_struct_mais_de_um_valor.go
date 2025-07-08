package main

import "fmt"

type pessoa struct {
	nome    string
	idade   int
	sabores []string
}

func main() {
	pessoa1 := pessoa{
		nome:    "Thiago",
		idade:   20,
		sabores: []string{"Morango", "Chocolate"},
	}

	fmt.Println(pessoa1)
} 