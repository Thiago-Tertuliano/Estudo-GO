package main

import "fmt"

func producer(out chan<- int) {
	defer close(out)
	for i := 0; i < 5; i++ {
		out <- i
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println(v)
	}
}