package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type Clube struct {
	Nome       string `json:"nome"`
	Modalidade string `json:"modalidade"`
	Fundacao   string `json:"fundacao"`
	Idade      int    `json:"idade"`
}

func main() {
	clubes := []Clube{
		{"São Paulo", "Futebol", "10/10/1930", 95},
		{"Corinthians", "Futebol", "10/10/1910", 115},
		{"Palmeiras", "Futebol", "10/10/1915", 110},
	}
	sort.Slice(clubes, func(i, j int) bool {
		return clubes[i].Nome < clubes[j].Nome
	})
	jsonData, err := json.MarshalIndent(clubes, "", "  ")
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
} 