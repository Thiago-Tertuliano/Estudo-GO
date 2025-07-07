package main

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Largura, Altura float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func mostrarArea(f Forma) {
	fmt.Printf("Área: %.2f\n", f.Area())
}

func main() {
	r := Retangulo{Largura: 10, Altura: 5}
	c := Circulo{Raio: 7}

	mostrarArea(r)
	mostrarArea(c)
} 