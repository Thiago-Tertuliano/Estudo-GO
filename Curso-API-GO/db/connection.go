package DB

import (
	"curso-api-go/configs"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func OpenConnection(config configs.DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
	
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com o banco: %w", err)
	}
	
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer ping no banco: %w", err)
	}
	
	return db, nil
}