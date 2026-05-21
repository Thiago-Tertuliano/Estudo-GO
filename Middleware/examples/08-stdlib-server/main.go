// Exemplo 08: servidor stdlib com mux, chain completa e rotas de teste.
package main

import (
	"fmt"
	"log"
	"net/http"

	"middleware-study/internal/httpmw"
)

func main() {
	mux := http.NewServeMux()

	api := httpmw.Chain(
		http.HandlerFunc(apiHandler),
		httpmw.Recover,
		httpmw.RequestID,
		httpmw.Logging,
	)
	mux.Handle("/api/", api)
	mux.Handle("GET /{$}", httpmw.Logging(http.HandlerFunc(home)))

	log.Println("listening :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Middleware study — rotas: /, /api/, /api/panic")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/panic":
		panic("erro interno simulado")
	case "/api/", "/api":
		id, _ := httpmw.GetRequestID(r.Context())
		fmt.Fprintf(w, "api ok request_id=%s\n", id)
	default:
		http.NotFound(w, r)
	}
}
