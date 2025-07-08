package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {
	pessoa1 := pessoa{
		nome:  "Thiago",
		idade: 20,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Thiago",
			idade: 20,
		},
		titulo:  "Desenvolvedor web",
		salario: 2000,
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa1.idade)
	fmt.Println(pessoa2)
	fmt.Println(pessoa2.salario)
} 