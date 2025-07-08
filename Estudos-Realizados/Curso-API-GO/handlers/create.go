package handlers

import (
	"curso-api-go/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

func Create(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
			return
		}
		
		var todo models.Todo
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, "erro ao decodificar JSON", http.StatusBadRequest)
			return
		}
		
		if todo.Title == "" {
			http.Error(w, "título é obrigatório", http.StatusBadRequest)
			return
		}
		
		id, err := models.Insert(db, todo)
		if err != nil {
			http.Error(w, "erro ao inserir todo", http.StatusInternalServerError)
			return
		}
		
		todo.ID = id
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(todo)
	}
}