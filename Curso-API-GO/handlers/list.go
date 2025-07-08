package handlers

import (
	"curso-api-go/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

func List(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
			return
		}
		
		todos, err := models.GetAll(db)
		if err != nil {
			http.Error(w, "erro ao buscar todos", http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	}
}