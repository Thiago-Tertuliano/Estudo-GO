// Exemplo 04: StatusRecorder captura o status HTTP real.
package main

import (
	"fmt"
	"log"
	"net/http"

	"middleware-study/internal/httpmw"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/created" {
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, "created")
			return
		}
		fmt.Fprint(w, "ok")
	})

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &httpmw.StatusRecorder{ResponseWriter: w, Status: http.StatusOK}
		handler.ServeHTTP(rec, r)
		log.Printf("status capturado: %d", rec.Status)
	}))

	log.Println("listening :8080 — teste / e /created")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
