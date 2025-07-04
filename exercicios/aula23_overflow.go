package main

import "fmt"

func main() {
	var i uint16
	i = 65535 // Valor máximo para uint16
	fmt.Println("Valor máximo de uint16:", i)

	i++ // Isso causará overflow, retornando ao valor mínimo
	fmt.Println("Após overflow, valor de uint16:", i) // Imprime 0

	i = 65536 // Tentando atribuir um valor maior que o máximo permitido
	fmt.Println("Após atribuição, valor de uint16:", i) // Imprime 0 novamente
} 