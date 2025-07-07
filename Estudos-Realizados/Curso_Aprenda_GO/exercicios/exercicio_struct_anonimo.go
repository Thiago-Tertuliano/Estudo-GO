package main

import "fmt"

func main() {
	x := struct {
		nome    string
		sabor   string
		ondetem []string
		vaibemcom map[string]string
	}{
		nome: "Bolo",
		sabor: "doce",
		ondetem: []string{"no brasil", "nos EUA"},
		vaibemcom: map[string]string{
			"no café da manhã": "café",
			"na café da tarde": "chá",
		},
	}

	fmt.Println(x)
} 