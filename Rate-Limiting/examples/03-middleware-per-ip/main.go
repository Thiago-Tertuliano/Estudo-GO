// Exemplo 03: rate limit por IP (RemoteAddr).
package main

import (
	"fmt"
	"log"
	"net/http"

	"rate-limiting-study/internal/ratelimit"
)

func main() {
	perIP := ratelimit.NewPerIP(2, 4)

	mux := http.NewServeMux()
	mux.Handle("GET /{$}", perIP.Middleware()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok from %s\n", r.RemoteAddr)
	})))

	log.Println("listening :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
