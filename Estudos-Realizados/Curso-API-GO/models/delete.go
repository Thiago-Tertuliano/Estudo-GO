package models

import (
	"database/sql"
	"errors"
)

func Delete(db *sql.DB, id int) error {
	query := `DELETE FROM todos WHERE id = $1`
	
	result, err := db.Exec(query, id)
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