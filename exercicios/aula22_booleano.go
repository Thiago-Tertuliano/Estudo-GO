package main

import "fmt"

var x bool
var a int
var y int
var z bool

func main() {
	fmt.Println(x)   // Imprime o valor inicial de x, que é false (valor zero para booleano)
	x = true         // Atribui o valor true à variável x
	fmt.Println(x)   // Imprime o novo valor de x, que agora é true
	x = (10 < 100)   // Atribui o resultado da expressão (10 < 100) a x, que será true
	fmt.Println(x)   // Imprime o valor de x, que ainda é true
	y := (10 == 100) // Declara uma nova variável y e atribui o resultado da expressão (10 == 100), que será false
 