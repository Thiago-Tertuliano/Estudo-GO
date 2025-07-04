package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func mudeMe(p *pessoa) {
	(*p).nome = "Caio"
	p.sobrenome = "Ribeiro"
}

func main() {
	thiago := pessoa{"Thiago", "Matos", 20}
	fmt.Println(thiago)
	mudeMe(&thiago)
	fmt.Println(thiago)
} 