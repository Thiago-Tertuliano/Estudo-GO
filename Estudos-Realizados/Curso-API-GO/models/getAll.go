package models

import "database/sql"

func GetAll(db *sql.DB) ([]Todo, error) {
	query := `
		SELECT id, title, description, done, created_at, updated_at
		FROM todos
		ORDER BY created_at DESC
	`
	
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.Done,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	
	return todos, nil
}