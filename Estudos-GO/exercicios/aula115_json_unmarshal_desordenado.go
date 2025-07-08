package main

import (
	"encoding/json"
	"fmt"
)

type Produto struct {
	Nome       string  `json:"nome"`
	Preco      float64 `json:"preco"`
	Quantidade int     `json:"quantidade"`
}

func main() {
	jsonRecebido := `{
	"quantidade": 3,
	"nome": "Lápis",
	"preco": 1.75
	}`
	var p Produto
	err := json.Unmarshal([]byte(jsonRecebido), &p)
	if err != nil {
		fmt.Println("Erro ao fazer unmarshal:", err)
		return
	}
	fmt.Println("Struct preenchida a partir do JSON:")
	fmt.Printf("Nome: %s\n", p.Nome)
	fmt.Printf("Preço: %.2f\n", p.Preco)
	fmt.Printf("Quantidade: %d\n", p.Quantidade)
} 