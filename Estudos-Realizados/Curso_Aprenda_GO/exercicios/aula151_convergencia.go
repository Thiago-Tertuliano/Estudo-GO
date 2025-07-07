package main

import (
	"fmt"
	"time"
)

func produtor(id int, out chan<- string) {
	for i := 0; i < 3; i++ {
		msg := fmt.Sprintf("Produtor %d: %d", id, i)
		out <- msg
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	saida := make(chan string)
	for i := 1; i <= 3; i++ {
		go produtor(i, saida)
	}
	for i := 0; i < 9; i++ {
		fmt.Println(<-saida)
	}
} 