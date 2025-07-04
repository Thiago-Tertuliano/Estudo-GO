package main

import (
	"fmt"
	"time"
	"sync"
)

var contador int
var mu sync.Mutex

func incrementar() {
	for i := 0; i < 1000; i++ {
		mu.Lock()
		contador++
		mu.Unlock()
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go incrementar()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Valor final do contador:", contador)
} 