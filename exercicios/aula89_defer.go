package main

import "fmt"

func main() {
	defer fmt.Println("com defer, veio primeiro")
	fmt.Println("sem defer, veio em segundo")
} 