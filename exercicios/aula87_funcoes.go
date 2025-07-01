package main

import "fmt"

func basica() {
	fmt.Println("Olá, bom dia!")
}

func soma(x, y int) int {
	return x + y
}

func somaVariavel(x ...int) (int, int) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x)
}

func main() {
	s := "abc"
	fmt.Println(len(s))
	basica()
	fmt.Println(soma(2, 3))
	fmt.Println(somaVariavel(1, 2, 3, 4))
} 