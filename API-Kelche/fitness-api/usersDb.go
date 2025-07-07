package repositories

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/storage"
	"os/user"
	"time"
)

func CreateUser(user models.User) (models.User, error) {
	db := storage.GetDB()
	sqlStatement := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password, time.Now(), time.Now()).Scan(&user.Id)
	if err != nil {
		return  models.User{}, err
	}

	user.CreatedAt = time.Now()
	user.UpdateAt = time.Now()
	return user, nil
}

func GetUsers() ([]models.User, error) {
	db := storage.GetDB()
	sqlStatement := `SELECT id, name, email, password, created_at, updated_at FROM users`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User 
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.password, &user.CreatedAt, &user.UptadeAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUser(id int) (models.User, error) {
	db := storage.GetDB()
	sqlStatement := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1`

	var user models.User
	err := db.QueryRow(sqlStatement, id).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdateAt)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func UpdateUser(user models.User, id int) (models.User, error) {
	db := storage.GetDB()
	sqlStatement := `UPDATE users SET name = $1, email = $2, password = $3, updated_at = $4 WHERE id = $5 RETURNING id`

	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password, time.Now()).Scan(&id)
	if err != nil {
		return models.User{}, err
	}

	user.Id = id
	user.UpdateAt = time.Now()
	return user, nil
}

func DeleteUser(id int) error {
	db := storage.GetDB()
	sqlStatement := `DELETE FROM users WHERE id = $1`

	_, err := db.Exec(sqlStatement, id)
	return err
}