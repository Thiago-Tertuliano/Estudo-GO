// Exemplo 07: recover de panic no handler.
package main

import (
	"fmt"
	"log"
	"net/http"

	"middleware-study/internal/httpmw"
)

func main() {
	mux := http.NewServeMux()
	h := httpmw.Chain(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			panic("simulacao de bug")
		}),
		httpmw.Recover,
		httpmw.Logging,
	)
	mux.Handle("GET /panic", h)
	mux.Handle("GET /{$}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok — tente GET /panic")
	}))

	log.Println("listening :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
