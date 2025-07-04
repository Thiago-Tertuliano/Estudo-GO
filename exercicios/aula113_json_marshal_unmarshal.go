package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func main() {
	p := Pessoa{
		Nome:  "Thiago",
		Idade: 30,
	}
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

	jsonRecebido := `{"nome":"Renata","idade":25}`
	var novaPessoa Pessoa
	err = json.Unmarshal([]byte(jsonRecebido), &novaPessoa)
	if err != nil {
		fmt.Println("Erro ao converter JSON para struct:", err)
		return
	}
	fmt.Println(novaPessoa)
} 