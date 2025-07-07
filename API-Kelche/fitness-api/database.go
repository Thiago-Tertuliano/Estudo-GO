package storage 

import (
	"database/sql" // O pacote "database/sql" fornece uma interface para acesso ao banco de dados
	"fmt" // O pacote "fmt" fornece funções para formatar e imprimir dados 
	"log" // O pacote "log" fornce uma interface para registrar mensagens de log
	"os" // O Pacote "os" fornece acesso ao sistema operacional

	_ "github.com/go-sql-driver/mysql" // O _ serve para ignorar o erro de importação
)

var db *sql.DB 

func GetDB() *sql.DB {
	return db
}

func ConnectDB() {
	var err error

	// Configurações do banco de dados
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "postgres")

	//String de conexão
	psqlInfo := fmt.Sprint("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname) //sslmode=disable desabilita o ssl que serve para segurança de conexão com o banco de dados.

	// Conexão ao banco de dados	
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal{"Erro ao conectar ao banco de dados:", err}
	}

	// Teste de conexão
	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	fmt.Println("Conectado ao banco de dados")
}

func getEnv(key, defaultValue) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
