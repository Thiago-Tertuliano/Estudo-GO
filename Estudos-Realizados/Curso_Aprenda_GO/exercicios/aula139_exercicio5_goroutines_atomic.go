package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var contador int32

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
			atomic.AddInt32(&contador, 1)
			v := atomic.LoadInt32(&contador)
			runtime.Gosched()
			fmt.Println(v)
			wg.Done()
		}()
	}
} 