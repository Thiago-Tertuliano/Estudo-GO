package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) nomecompleto() {
	fmt.Println(p.nome, p.sobrenome)
}

func main() {
	Thiago := pessoa{"Thiago", "Tertuliano", 20}
	Thiago.nomecompleto()
} 