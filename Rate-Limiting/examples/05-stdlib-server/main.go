// Exemplo 05: servidor com chain — recover simples + rate limit global.
package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/time/rate"

	"rate-limiting-study/internal/ratelimit"
)

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v", err)
				http.Error(w, "internal error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	limiter := rate.NewLimiter(3, 6)

	mux := http.NewServeMux()
	api := ratelimit.Chain(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case "/api/panic":
				panic("simulated")
			default:
				fmt.Fprintln(w, "api ok")
			}
		}),
		recoverMiddleware,
		ratelimit.Middleware(limiter),
	)
	mux.Handle("/api/", api)
	mux.Handle("GET /{$}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "home — try /api/ and rapid curl for 429")
	}))

	log.Println("listening :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
