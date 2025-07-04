package main

import "fmt"

const (
	anoAtual = 2025
	proximoAno = iota + anoAtual
	segundoAno = iota + anoAtual
	terceiroAno = iota + anoAtual
	quartoAno = iota + anoAtual
)

func main() {
	fmt.Println(proximoAno, segundoAno, terceiroAno, quartoAno)
} 