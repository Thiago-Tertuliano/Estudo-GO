package main

import "fmt"

func main() {
	x := struct {
		nome  string
		idade int
	}{
		nome:  "Mario",
		idade: 20,
	}

	fmt.Println(x)
} 