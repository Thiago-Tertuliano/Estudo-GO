package main

import "fmt"

func main() {
	estados := []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espirito Santo", "Goiás", "Maranhão", "Mato Grosso do Sul", "Mato Grosso", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(len(estados), cap(estados))
	fmt.Println(estados)
} 