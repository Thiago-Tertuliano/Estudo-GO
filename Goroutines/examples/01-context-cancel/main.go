package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d: parou (%v)\n", id, ctx.Err())
			return
		default:
			// Simula trabalho; em produção seria I/O com ctx
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("worker %d: tick\n", id)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go worker(ctx, 1)

	<-ctx.Done()
	fmt.Println("main: context cancelado, aguardando worker...")
	time.Sleep(300 * time.Millisecond)
}