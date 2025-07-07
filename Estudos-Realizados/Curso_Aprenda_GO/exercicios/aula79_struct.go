package main

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	cliente1 := cliente{
		nome:      "Thiago",
		sobrenome: "Matos",
		fumante:   false,
	}

	fmt.Println(cliente1)
} 