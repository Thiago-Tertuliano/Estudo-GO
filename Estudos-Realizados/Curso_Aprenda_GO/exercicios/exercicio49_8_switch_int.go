package main

import "fmt"

func main() {
	var idade int = 14

	switch idade {
	case 14:
		fmt.Printf("O jogador é sub-14")
	case 16:
		fmt.Printf("O jogador é sub-16")
	case 18:
		fmt.Printf("O jogador é sub-18")
	case 20:
		fmt.Printf("O jogador é sub-20")
	case 21:
		fmt.Printf("O jogador é Profissional")
	default:
		fmt.Printf("Ele não é jogador!")
	}

	switch idade {
	case 14, 16, 18, 20, 21:
		fmt.Printf("\nEle é jogador!")
	default:
		fmt.Printf("\nEle não é jogador!")
	}
} 