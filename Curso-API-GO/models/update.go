package models

import (
	"database/sql"
	"errors"
)

func Update(db *sql.DB, id int, todo Todo) error {
	query := `
		UPDATE todos
		SET title = $1, description = $2, done = $3, updated_at = NOW()
		WHERE id = $4
	`
	
	result, err := db.Exec(query, todo.Title, todo.Description, todo.Done, id)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return errors.New("todo não encontrado")
	}
	
	return nil
}