package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type User struct {
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Idade     int    `json:"idade"`
}

func main() {
	usuarios := []User{
		{"Ana", "Silva", 30},
		{"Carlos", "Oliveira", 25},
		{"Bruno", "Almeida", 25},
		{"Daniela", "Silva", 40},
		{"Ana", "Costa", 30},
	}
	sort.Slice(usuarios, func(i, j int) bool {
		if usuarios[i].Idade == usuarios[j].Idade {
			return usuarios[i].Sobrenome < usuarios[j].Sobrenome
		}
		return usuarios[i].Idade < usuarios[j].Idade
	})
	fmt.Println("Usuários ordenados por idade e sobrenome:")
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(usuarios)
	if err != nil {
		fmt.Println("Erro ao codificar JSON:", err)
	}
} 