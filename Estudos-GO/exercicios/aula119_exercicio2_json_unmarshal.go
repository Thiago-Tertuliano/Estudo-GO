package main

import (
	"encoding/json"
	"fmt"
)

type Jogador struct {
	Nome         string `json:"nome"`
	NumeroCamisa int    `json:"numeroCamisa"`
	Posicao      string `json:"posicao"`
}

func main() {
	jsonRecebido := `{
	"nome": "Thiago",
	"numeroCamisa": 10,
	"posicao": "Atacante"
	}`
	var j Jogador
	err := json.Unmarshal([]byte(jsonRecebido), &j)
	if err != nil {
		fmt.Println("Erro ao fazer o unmarshal:", err)
		return
	}
	fmt.Println("Struct preenchida a partir do JSON:")
	fmt.Printf("Nome do jogador: %s\n", j.Nome)
	fmt.Printf("Número da camisa: %d\n", j.NumeroCamisa)
	fmt.Printf("Posição: %s\n", j.Posicao)
} 