package handlers

import (
	"curso-api-go/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
)

func Delete(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
			return
		}
		
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}
		
		err = models.Delete(db, id)
		if err != nil {
			if err.Error() == "todo não encontrado" {
				http.Error(w, "todo não encontrado", http.StatusNotFound)
				return
			}
			http.Error(w, "erro ao deletar todo", http.StatusInternalServerError)
			return
		}
		
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "todo deletado com sucesso"})
	}
}