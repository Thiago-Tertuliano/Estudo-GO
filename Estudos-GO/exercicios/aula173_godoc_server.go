// Pacote saudacao fornece funções para gerar mensagens de boas-vindas.
package main

import "fmt"

// BemVindo retorna uma mensagem de boas-vindas para o nome fornecido.
func BemVindo(nome string) string {
	return "Bem-vindo, " + nome + "!"
}

func main() {
	fmt.Println(BemVindo("Thiago"))
} 