// Exemplo 05: cadeia de middlewares com httpmw.Chain.
package main

import (
	"fmt"
	"log"
	"net/http"

	"middleware-study/internal/httpmw"
)

func withHeader(name, value string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(name, value)
			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	final := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "chain ok")
	})
	h := httpmw.Chain(final,
		httpmw.Recover,
		httpmw.Logging,
		withHeader("X-Study", "middleware-chain"),
	)

	mux := http.NewServeMux()
	mux.Handle("GET /{$}", h)
	log.Println("listening :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
