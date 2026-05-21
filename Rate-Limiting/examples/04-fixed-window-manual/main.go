// Exemplo 04: janela fixa manual (5 req a cada 10s por IP).
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"rate-limiting-study/internal/ratelimit"
)

func main() {
	fw := ratelimit.NewFixedWindow(5, 10*time.Second)

	mux := http.NewServeMux()
	mux.Handle("GET /{$}", fw.Middleware(func(r *http.Request) string {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			return r.RemoteAddr
		}
		return ip
	})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})))

	log.Println("listening :8080 — max 5 req / 10s por IP")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
