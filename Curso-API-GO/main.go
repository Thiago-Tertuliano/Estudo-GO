package main

import (
	"curso-api-go/configs"
	"curso-api-go/DB"
	"curso-api-go/handlers"
	"fmt"
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Carregar configurações
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Erro ao carregar configurações:", err)
	}
	
	// Conectar ao banco de dados
	db, err := DB.OpenConnection(config.GetDBConfig())
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()
	
	// Configurar rotas
	r := chi.NewRouter()
	
	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Timeout(60))
	
	// Rotas da API
	r.Route("/api", func(r chi.Router) {
		r.Route("/todos", func(r chi.Router) {
			r.Get("/", handlers.List(db))           // GET /api/todos
			r.Post("/", handlers.Create(db))        // POST /api/todos
			r.Get("/{id}", handlers.Get(db))        // GET /api/todos/{id}
			r.Put("/{id}", handlers.Update(db))     // PUT /api/todos/{id}
			r.Delete("/{id}", handlers.Delete(db))  // DELETE /api/todos/{id}
		})
	})
	
	// Rota de health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API está funcionando!"))
	})
	
	// Iniciar servidor
	port := config.GetAPIConfig().Port
	fmt.Printf("Servidor iniciado na porta %s\n", port)
	fmt.Println("Endpoints disponíveis:")
	fmt.Println("- GET    /api/todos")
	fmt.Println("- POST   /api/todos")
	fmt.Println("- GET    /api/todos/{id}")
	fmt.Println("- PUT    /api/todos/{id}")
	fmt.Println("- DELETE /api/todos/{id}")
	fmt.Println("- GET    /health")
	
	log.Fatal(http.ListenAndServe(port, r))
}