package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "tarde demais"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("timeout")
	}
}