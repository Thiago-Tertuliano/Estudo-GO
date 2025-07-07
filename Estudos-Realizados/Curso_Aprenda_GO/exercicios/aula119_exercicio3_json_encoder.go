package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Jogador struct {
	Nome         string `json:"nome"`
	NumeroCamisa int    `json:"numeroCamisa"`
	Posicao      string `json:"posicao"`
}

func main() {
	j := Jogador{
		Nome:         "Thiago",
		NumeroCamisa: 10,
		Posicao:      "Atacante",
	}
	fmt.Println("JSON gerado com json.NewEncoder(os.Stdout):")
	encoder := json.NewEncoder(os.Stdout)
	err := encoder.Encode(j)
	if err != nil {
		fmt.Println("Erro ao codificar JSON:", err)
	}
} 