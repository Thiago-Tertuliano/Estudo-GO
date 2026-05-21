// Exemplo 06: request ID no context e no header de resposta.
package main

import (
	"fmt"
	"log"
	"net/http"

	"middleware-study/internal/httpmw"
)

func main() {
	final := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, ok := httpmw.GetRequestID(r.Context())
		if !ok {
			http.Error(w, "sem request id", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "request_id=%s\n", id)
	})
	h := httpmw.Chain(final, httpmw.RequestID, httpmw.Logging)

	mux := http.NewServeMux()
	mux.Handle("GET /{$}", h)
	log.Println("listening :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
