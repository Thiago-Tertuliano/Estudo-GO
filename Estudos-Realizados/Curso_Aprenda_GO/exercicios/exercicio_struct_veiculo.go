package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	carrao := sedan{veiculo{4, "laranja"}, true}
	fuscavo := caminhonete{veiculo{2, "Azul"}, true}

	fmt.Println(carrao)
	fmt.Println(fuscavo)
} 