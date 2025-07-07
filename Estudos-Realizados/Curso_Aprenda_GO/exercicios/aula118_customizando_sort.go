package main

import (
	"fmt"
	"sort"
)

type Pessoa struct {
	Nome   string
	Idade  int
	Cidade string
}

func main() {
	pessoas := []Pessoa{
		{"João", 30, "São Paulo"},
		{"Maria", 25, "Rio de Janeiro"},
		{"Pedro", 25, "Belo Horizonte"},
		{"Ana", 40, "Curitiba"},
	}
	sort.Slice(pessoas, func(i, j int) bool {
		if pessoas[i].Idade == pessoas[j].Idade {
			return pessoas[i].Nome < pessoas[j].Nome
		}
		return pessoas[i].Idade < pessoas[j].Idade
	})
	fmt.Println("Pessoas ordenadas por idade e nome:")
	for _, p := range pessoas {
		fmt.Printf("%s - %d anos - %s\n", p.Nome, p.Idade, p.Cidade)
	}
} 