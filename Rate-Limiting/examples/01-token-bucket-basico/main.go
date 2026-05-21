// Exemplo 01: rate.Limiter sem HTTP — Allow() e rejeição.
package main

import (
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	// 2 por segundo, burst 3
	l := rate.NewLimiter(2, 3)

	for i := 1; i <= 8; i++ {
		if l.Allow() {
			fmt.Printf("req %d: permitida\n", i)
		} else {
			fmt.Printf("req %d: negada (sem token)\n", i)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
