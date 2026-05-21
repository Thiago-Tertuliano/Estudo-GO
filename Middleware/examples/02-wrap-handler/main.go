// Exemplo 02: primeiro middleware — wrap manual de http.Handler.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func withBanner(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("->", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Println("<-", r.Method, r.URL.Path)
	})
}

func main() {
	final := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "handler final")
	})
	handler := withBanner(final)

	mux := http.NewServeMux()
	mux.Handle("GET /{$}", handler)
	log.Println("listening :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
