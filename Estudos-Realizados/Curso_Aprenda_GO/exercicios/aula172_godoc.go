// Pacote calculadora fornece funções matemáticas básicas.
package main

import "fmt"

// Soma retorna a soma de dois inteiros.
func Soma(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(Soma(2, 3))
} 