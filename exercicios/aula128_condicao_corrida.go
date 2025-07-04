package main

import (
	"fmt"
	"time"
)

var contador int

func incrementar() {
	for i := 0; i < 1000; i++ {
		contador++
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go incrementar()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Valor final do contador:", contador)
} 