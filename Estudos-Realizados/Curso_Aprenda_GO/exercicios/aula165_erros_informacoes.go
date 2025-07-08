package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("registro não encontrado")

func buscarRegistro(id int) error {
	if id == 0 {
		return fmt.Errorf("ID inválido (%d): %w", id, ErrNotFound)
	}
	return nil
}

func main() {
	err := buscarRegistro(0)
	if err != nil {
		fmt.Println("Erro:", err)
		if errors.Is(err, ErrNotFound) {
			fmt.Println("Tratamento específico para registro não encontrado")
		}
	}
} 