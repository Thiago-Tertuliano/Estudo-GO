package main

import "fmt"

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("Número: %d, Caractere: %c, Binário: %b, Decimal: %d, Hexadecimal: %#x\n", i, i, i, i, i)
	}
} 