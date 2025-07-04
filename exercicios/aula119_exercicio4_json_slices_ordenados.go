package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type DadosOrdenados struct {
	Numeros []int    `json:"numeros"`
	Nomes   []string `json:"nomes"`
}

func main() {
	numeros := []int{42, 7, 19, 3, 99}
	nomes := []string{"Carlos", "Ana", "Bruno", "Daniela"}
	sort.Ints(numeros)
	sort.Strings(nomes)
	dados := DadosOrdenados{
		Numeros: numeros,
		Nomes:   nomes,
	}
	fmt.Println("JSON com slices ordenados:")
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(dados)
	if err != nil {
		fmt.Println("Erro ao codificar JSON:", err)
	}
} 