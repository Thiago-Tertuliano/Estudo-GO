package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	texto := "Olá, mundo! Este é um exemplo da interface Writer."
	var escritor io.Writer = os.Stdout
	_, err := escritor.Write([]byte(texto))
	if err != nil {
		fmt.Println("Erro ao escrever:", err)
	}
	fmt.Println("\n---")
	var buffer strings.Builder
	_, err = buffer.Write([]byte("Texto armazenado no buffer."))
	if err != nil {
		fmt.Println("Erro ao escrever no buffer:", err)
	}
	fmt.Println(buffer.String())
} 