package main

import (
	"errors"
	"fmt"
	"log"
)

func retornaErro() error {
	return errors.New("algo deu errado")
}

func main() {
	err := retornaErro()
	if err != nil {
		fmt.Println("Usando fmt:", err)
		log.Println("Usando log:", err)
	}
} 