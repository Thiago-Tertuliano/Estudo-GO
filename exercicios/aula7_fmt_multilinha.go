package main

import "fmt"

func main() {
	x := `O acento grave é usado para \n criar\n\n\n strings multilinha ou para evitar %v a necessidade\n\nde escapar de caracteres especiais.`
	fmt.Println(x)
} 