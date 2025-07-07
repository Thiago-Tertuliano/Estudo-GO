package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var contador int

const quantidadeDeGoroutines = 50000

func main() {
	criarGoroutines(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Total de goroutines:\t", quantidadeDeGoroutines, "\nTotal do contador:\t", contador)
}

func criarGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}
} 