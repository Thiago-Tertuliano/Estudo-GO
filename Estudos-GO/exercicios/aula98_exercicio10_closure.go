package main

import "fmt"

func contador() func() int {
	contadorExterno := 0
	return func() int {
		contadorExterno++
		return contadorExterno
	}
}

func main() {
	cont1 := contador()
	cont2 := contador()

	fmt.Println(cont1())
	fmt.Println(cont2())
} 