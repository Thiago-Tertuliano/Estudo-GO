// Exemplo 03: middleware de logging (método, path, status, duração).
package main

import (
	"fmt"
	"log"
	"net/http"

	"middleware-study/internal/httpmw"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /{$}", httpmw.Logging(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})))
	mux.Handle("GET /error", httpmw.Logging(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "bad", http.StatusBadRequest)
	})))
	log.Println("listening :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
