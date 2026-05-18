package main

import (
	"fmt"
	"sync"
)

func generate(out chan<- int, n int) {
	defer close(out)
	for i := 1; i <= n; i++ {
		out <- i
	}
}

func process(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)
	for v := range in {
		out <- v * 2
	}
}

func aggregate(in <-chan int) int {
	sum := 0
	for v := range in {
		sum += v
	}
	return sum
}

func main() {
	const n = 5
	raw := make(chan int, 2)
	mid := make(chan int, 2)

	var wg sync.WaitGroup
	wg.Add(1)
	go generate(raw, n)
	go process(raw, mid, &wg)

	sum := aggregate(mid)
	wg.Wait()

	fmt.Println("soma (1.5)*2 =", sum) // 2+4+6+8+10 = 30
}