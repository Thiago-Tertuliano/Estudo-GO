package models

import (
	"database/sql"
	"errors"
)

func Get(db *sql.DB, id int) (Todo, error) {
	query := `
		SELECT id, title, description, done, created_at, updated_at
		FROM todos
		WHERE id = $1
	`
	
	var todo Todo
	err := db.QueryRow(query, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.Done,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return Todo{}, errors.New("todo não encontrado")
		}
		return Todo{}, err
	}
	
	return todo, nil
}