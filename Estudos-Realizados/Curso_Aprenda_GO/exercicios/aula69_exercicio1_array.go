package main

import "fmt"

func main() {
	val := [5]int{1, 2, 3, 4, 5}

	for i, v := range val {
		fmt.Println(i, v)
	}

	fmt.Printf("%T, %v", val, val)
} 