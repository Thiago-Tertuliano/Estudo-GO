package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d: shutdown\n", id)
			return
		case <-time.After(400 * time.Millisecond):
			fmt.Printf("worker %d: trabalhando...\n", id)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	fmt.Println("rodando... Ctrl+C para encerrar")
	<-sigCh
	fmt.Println("recebido sinal de encerramento, cancelando contextos...")
	cancel()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("shutdown completo")
	case <-time.After(3 * time.Second):
		fmt.Println("timeout: encerrando sem shutdown completo")
	}
}