// Exemplo 09: mesmo middleware stdlib reutilizado com chi.Use.
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"middleware-study/internal/httpmw"
)

func main() {
	r := chi.NewRouter()
	r.Use(httpmw.Recover)
	r.Use(httpmw.RequestID)
	r.Use(httpmw.Logging)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "chi + httpmw")
	})
	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("chi panic test")
	})

	log.Println("chi bridge listening :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
