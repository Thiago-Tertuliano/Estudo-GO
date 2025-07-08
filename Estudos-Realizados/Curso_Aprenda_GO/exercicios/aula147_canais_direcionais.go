package main

import "fmt"

func enviar(ch chan<- string) {
	ch <- "Olá de enviar()"
}

func receber(ch <-chan string) {
	msg := <-ch
	fmt.Println("Recebido:", msg)
}

func main() {
	ch := make(chan string)
	go enviar(ch)
	receber(ch)
} 