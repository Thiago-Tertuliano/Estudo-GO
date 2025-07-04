package main

import "fmt"

func main() {
	mepe := map[string][]string{
		"dasilva_guriaestranhadostrangerthings": {
			"desaparecer", "ser assustadora", "raspar o cabelo",
		},
		"senna_ayrton": {
			"andar de jetski", "ser milionario", "ser paulista",
		},
		"ronaldo_cristiano": {
			"melhor do mundo", "bilionario", "jogador do real madrid",
		},
	}

	for k, v := range mepe {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}

	fmt.Println(mepe)
} 