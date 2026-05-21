// Exemplo 02: um rate.Limiter global para todas as rotas.
package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/time/rate"

	"rate-limiting-study/internal/ratelimit"
)

func main() {
	// Limite baixo para testar 429 com curl em loop
	l := rate.NewLimiter(1, 2)

	mux := http.NewServeMux()
	mux.Handle("GET /{$}", ratelimit.Middleware(l)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})))

	log.Println("listening :8080 — envie várias requisições rápidas")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
