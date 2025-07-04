package main

import "fmt"

func main() {
	sabores := []string{"Peperoni", "Mussarela", "Abacaxi", "4 Queijos", "Margherita"}

	fatia := sabores[:]

	fmt.Println(fatia)
	fmt.Println(sabores[0], sabores[1], sabores[2], sabores[3], sabores[4])
	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}
} 