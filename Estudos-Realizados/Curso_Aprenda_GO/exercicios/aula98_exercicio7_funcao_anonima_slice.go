package main

import "fmt"

func main() {
	numeros := []int{10, 20}

	soma := func(nums []int) int {
		total := 0
		for _, n := range nums {
			total += n
		}
		return total
	}

	resultado := soma(numeros)
	fmt.Printf("A soma é: %d\n", resultado)
} 