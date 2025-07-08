package main

import "fmt"

func sayHello(ch chan string) {
	ch <- "Olá de uma goroutine!"
}

func main() {
	ch := make(chan string)
	go sayHello(ch)
	msg := <-ch
	fmt.Println(msg)
} 