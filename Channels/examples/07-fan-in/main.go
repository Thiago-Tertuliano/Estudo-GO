package main

import (
	"fmt"
	"sync"
)

func producer(name string, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		out <- fmt.Sprintf("%s-%d", name, i)
	}
}

func merge(out chan<- string, wg *sync.WaitGroup, ins ...<-chan string) {
	defer close(out)
	var w sync.WaitGroup
	multiplex := func(c <-chan string) {
		defer w.Done()
		for v := range c {
			out <- v
		}
	}
	w.Add(len(ins))
	for _, c := range ins {
		go multiplex(c)
	}
	w.Wait()
	wg.Done()
}

func main() {
	a, b := make(chan string), make(chan string)
	var prodWG sync.WaitGroup
	prodWG.Add(2)
	go producer("A", a, &prodWG)
	go producer("B", b, &prodWG)
	go func() {
		prodWG.Wait()
		close(a)
		close(b)
	}()

	merged := make(chan string)
	var mergeWG sync.WaitGroup
	mergeWG.Add(1)
	go merge(merged, &mergeWG, a, b)

	for v := range merged {
		fmt.Println(v)
	}
}