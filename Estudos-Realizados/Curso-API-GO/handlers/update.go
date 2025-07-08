package handlers

import (
	"curso-api-go/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
)

func Update(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
			return
		}
		
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}
		
		var todo models.Todo
		err = json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, "erro ao decodificar JSON", http.StatusBadRequest)
			return
		}
		
		if todo.Title == "" {
			http.Error(w, "título é obrigatório", http.StatusBadRequest)
			return
		}
		
		err = models.Update(db, id, todo)
		if err != nil {
			if err.Error() == "todo não encontrado" {
				http.Error(w, "todo não encontrado", http.StatusNotFound)
				return
			}
			http.Error(w, "erro ao atualizar todo", http.StatusInternalServerError)
			return
		}
		
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "todo atualizado com sucesso"})
	}
}