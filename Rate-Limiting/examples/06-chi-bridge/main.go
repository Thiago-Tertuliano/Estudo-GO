// Exemplo 06: Chi com ratelimit.Middleware (assinatura stdlib).
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"golang.org/x/time/rate"

	"rate-limiting-study/internal/ratelimit"
)

func main() {
	r := chi.NewRouter()
	l := rate.NewLimiter(2, 4)
	r.Use(ratelimit.Middleware(l))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "chi + rate limit")
	})

	log.Println("chi bridge listening :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
