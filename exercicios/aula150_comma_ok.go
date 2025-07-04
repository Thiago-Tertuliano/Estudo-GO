package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 10
		close(ch)
	}()
	val, ok := <-ch
	fmt.Println("Valor:", val, "OK:", ok)
	val, ok = <-ch
	fmt.Println("Valor:", val, "OK:", ok)
} 