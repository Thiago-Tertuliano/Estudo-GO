package main

import "fmt"

func main() {
	x := `O acento grave é usado para \n criar\n\n\n strings multilinha ou para evitar %v a necessidade\n\nde escapar de caracteres especiais.`
	fmt.Println(x)

	y := "3Con"
	z := " IT & Digital"
	w := fmt.Sprint(y, z)
	fmt.Println(w)

	a := 42
	b := "James Bond"
	c := true
	fmt.Println(a, b, c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	s := fmt.Sprintf("%v, %v, %v", a, b, c)
	fmt.Println(s)
} 