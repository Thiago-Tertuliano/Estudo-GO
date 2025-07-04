package main

import "fmt"

func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}

func main() {
	num := 5
	fmt.Printf("Fatorial de %d é %d\n", num, fatorial(num))
} 