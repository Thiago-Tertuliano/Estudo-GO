package main

import "fmt"

func soma() int {
	return 25
}

func somador() (int, string) {
	return 10, "dez"
}

func main() {
	fmt.Println(soma())
	fmt.Println(somador())
} 