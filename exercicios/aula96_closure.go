package main

import "fmt"

func contador() func() int {
	contadorInterno := 0
	return func() int {
		contadorInterno++
		return contadorInterno
	}
}

func main() {
	cont1 := contador()
	cont2 := contador()

	fmt.Println(cont1())
	fmt.Println(cont1())
	fmt.Println(cont2())
	fmt.Println(cont1())
	fmt.Println(cont2())
} 