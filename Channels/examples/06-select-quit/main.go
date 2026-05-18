package main

import "fmt"

func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func main() {
	done := make(chan struct{})
	defer close(done)

	c := gen(done, 1, 2, 3, 4, 5)
	for v := range c {
		if v == 3 {
			break // em app real: cancel context
		}
		fmt.Println(v)
	}
}