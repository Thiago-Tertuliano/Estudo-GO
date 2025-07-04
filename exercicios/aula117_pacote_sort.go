package main

import (
	"fmt"
	"sort"
)

func main() {
	numeros := []int{42, 7, 19, 3, 99}
	sort.Ints(numeros)
	fmt.Println("Slice ordenado:", numeros)

	nomes := []string{"Carlos", "Ana", "Bruno", "Daniela"}
	sort.Strings(nomes)
	fmt.Println("Nomes ordenados:", nomes)

	type Pessoa struct {
		Nome string
		Idade int
	}

	pessoas := []Pessoa{
		{"João", 30},
		{"Maria", 25},
		{"Pedro", 40},
	}
	sort.Slice(pessoas, func(i, j int) bool {
		return pessoas[i].Idade < pessoas[j].Idade
	})
	fmt.Println("Pessoas ordenadas por idade:")
	for _, p := range pessoas {
		fmt.Printf("%s - %d anos\n", p.Nome, p.Idade)
	}
} 