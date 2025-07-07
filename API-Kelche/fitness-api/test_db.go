package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	psqlInfo := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Erro ao fazer o ping no banco de dados:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao fazer o ping no banco de dados:", err)
	}

	fmt.Println("Conectado ao banco de dados")

	// Teste se as tabelas existem

	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		log.Fatal("Erro ao consultar as tabelas:", err)
	}
	defer rows.Close()

	fmt.Println("Tabelas existentes:")
	for rows.Next() {
		var tableName string 
		err := rows.Scan(&tableName)
		if err != nil {
			log.Fatal("Erro ao ler nome da tabela:", err)
		}
		fmt.Println(" -", tableName)
	}
}