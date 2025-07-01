package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("arquivo_inexistente.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()
} 