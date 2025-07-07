package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type Produto struct {
	Nome      string  `json:"nome"`
	Preco     float64 `json:"preco"`
	Quantidade int    `json:"quantidade"`
}

func main() {
	produtos := []Produto{
		{"Caderno", 12.90, 5},
		{"Caneta", 2.50, 10},
		{"Borracha", 1.20, 8},
	}
	sort.Slice(produtos, func(i, j int) bool {
		return produtos[i].Nome < produtos[j].Nome
	})
	jsonData, err := json.MarshalIndent(produtos, "", "  ")
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
} 