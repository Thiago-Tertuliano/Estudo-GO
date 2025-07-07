package main

import "fmt"

func main() {
	info := [][]string{
		{"Thiago"},
		{"Tertuliano"},
		{"Jogar Minecraft"},
	}
	for _, v := range info {
		fmt.Println(v[0])
		for _, dados := range v {
			fmt.Println("\t", dados)
		}
	}
} 