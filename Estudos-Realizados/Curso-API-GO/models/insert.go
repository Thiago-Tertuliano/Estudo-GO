package models

import (
	"database/sql"
)

func Insert(db *sql.DB, todo Todo) (int, error) {
	query := `
		INSERT INTO todos (title, description, done, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id
	`
	
	var id int
	err := db.QueryRow(query, todo.Title, todo.Description, todo.Done).Scan(&id)
	if err != nil {
		return 0, err
	}
	
	return id, nil
}