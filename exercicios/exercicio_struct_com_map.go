package main

import "fmt"

type pessoa struct {
	nome    string
	idade   int
	sabores []string
}

func main() {
	meumapa := make(map[string]pessoa)

	meumapa["Renata"] = pessoa{
		nome:    "Renata",
		idade:   34,
		sabores: []string{"Chocolate", "Morango"},
	}

	for _, valor := range meumapa {
		fmt.Println(valor)
	}
} 