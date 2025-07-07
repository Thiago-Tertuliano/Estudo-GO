package main

import "fmt"

func dobrarNumero(n *int) {
	*n = *n * 2
}

func main() {
	num := 10
	fmt.Println("Antes:", num)
	dobrarNumero(&num)
	fmt.Println("Depois:", num)
} 