// Exemplo 01: handler HTTP mínimo sem middleware.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	})
	log.Println("listening :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
