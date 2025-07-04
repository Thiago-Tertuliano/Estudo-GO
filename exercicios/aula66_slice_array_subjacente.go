package main

import "fmt"

func main() {
	primeiroSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(primeiroSlice)

	segundoSlice := append(primeiroSlice[:2], primeiroSlice[4:]...)

	fmt.Println(segundoSlice)
	fmt.Println(primeiroSlice)
} 