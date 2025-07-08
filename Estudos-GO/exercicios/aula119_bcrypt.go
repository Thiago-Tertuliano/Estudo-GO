package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "minhaSenha123"
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Erro ao gerar hash:", err)
		return
	}
	fmt.Println("Hash gerado:", string(hash))

	err = bcrypt.CompareHashAndPassword(hash, []byte(senha))
	if err != nil {
		fmt.Println("Senha incorreta!")
	} else {
		fmt.Println("Senha correta!")
	}
} 