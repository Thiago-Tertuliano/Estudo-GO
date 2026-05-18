package main

import "fmt"

func worker(out chan<- string) {
	out <- "ok"
}

func main() {
	ch := make(chan string)
	go worker(ch)
	fmt.Println(<-ch)
}