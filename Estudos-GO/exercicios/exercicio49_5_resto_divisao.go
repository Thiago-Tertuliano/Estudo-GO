package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Println("Este número é divisivel por 4: ", i)
		}
	}
} 