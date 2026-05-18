package main

import "fmt"

func main() {
	jobs := make(chan int, 3)
	jobs <- 1
	jobs <- 2
	jobs <- 3
	fmt.Println("Buffer cheio; próximo send bloquearia até ler")
	fmt.Println(<-jobs, <-jobs, <-jobs)
}