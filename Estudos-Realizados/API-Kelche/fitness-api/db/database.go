package storage 

/*
Este arquivo (database.go) é responsável por toda a configuração e conexão com o banco de dados.
É aqui que definimos como nossa aplicação se conecta ao PostgreSQL.

Conceitos importantes:
- Package storage: agrupa funcionalidades relacionadas ao armazenamento de dados
- Variável global db: mantém a conexão ativa com o banco durante toda a execução
- Função ConnectDB(): estabelece a conexão inicial
- Função GetDB(): permite que outros pacotes acessem a conexão
- Variáveis de ambiente: permitem configurar o banco sem alterar o código
*/

import (
	"database/sql" // Interface padrão do Go para trabalhar com bancos de dados SQL
	"fmt"          // Para formatação de strings (como printf em C)
	"log"          // Para registrar mensagens de log (erros, informações)
	"os"           // Para acessar variáveis de ambiente do sistema operacional

	_ "github.com/lib/pq" // Driver PostgreSQL para Go (o _ ignora o erro de importação)
)

// Variável global que armazena a conexão com o banco de dados
// É global para que possa ser acessada por outros pacotes
var db *sql.DB 

// Função GetDB() retorna a conexão com o banco de dados
// Outros pacotes usam esta função para executar queries
func GetDB() *sql.DB {
	return db
}

// Função ConnectDB() estabelece a conexão com o banco de dados PostgreSQL
// Esta função é chamada no início da aplicação (main.go)
func ConnectDB() {
	var err error // Variável para armazenar erros

	// ===== CONFIGURAÇÃO DO BANCO DE DADOS =====
	// Usamos variáveis de ambiente para configurar o banco
	// Isso permite mudar a configuração sem alterar o código
	// getEnv() retorna o valor da variável de ambiente ou um valor padrão
	
	host := getEnv("DB_HOST", "localhost")     // Endereço do servidor PostgreSQL
	port := getEnv("DB_PORT", "5432")          // Porta padrão do PostgreSQL
	user := getEnv("DB_USER", "postgres")      // Usuário do banco
	password := getEnv("DB_PASSWORD", "123456") // Senha do usuário
	dbname := getEnv("DB_NAME", "postgres")    // Nome do banco de dados

	// ===== STRING DE CONEXÃO =====
	// PostgreSQL usa uma string específica para conexão
	// fmt.Sprintf() formata a string substituindo %s pelos valores
	// sslmode=disable desabilita SSL (apenas para desenvolvimento)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		host, port, user, password, dbname)

	// ===== ESTABELECENDO A CONEXÃO =====
	// sql.Open() cria uma conexão com o banco
	// "postgres" = nome do driver
	// psqlInfo = string de conexão
	// Retorna *sql.DB (ponteiro para a conexão) e error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		// log.Fatal() para a aplicação e exibe a mensagem de erro
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	// ===== TESTE DE CONEXÃO =====
	// db.Ping() testa se a conexão está funcionando
	// É importante testar antes de usar a conexão
	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	// Mensagem de sucesso
	fmt.Println("✅ Conectado ao banco de dados PostgreSQL com sucesso!")
}

// Função getEnv() obtém o valor de uma variável de ambiente
// Se a variável não existir, retorna o valor padrão
// key = nome da variável de ambiente
// defaultValue = valor a retornar se a variável não existir
func getEnv(key, defaultValue string) string {
	// os.Getenv() retorna o valor da variável de ambiente
	// Se a variável não existir, retorna string vazia ""
	if value := os.Getenv(key); value != "" {
		return value // Retorna o valor da variável de ambiente
	}
	return defaultValue // Retorna o valor padrão
}

/*
EXEMPLO DE USO DAS VARIÁVEIS DE AMBIENTE:

No Windows (CMD):
set DB_HOST=localhost
set DB_PORT=5432
set DB_USER=meu_usuario
set DB_PASSWORD=minha_senha
set DB_NAME=meu_banco

No Linux/Mac:
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=meu_usuario
export DB_PASSWORD=minha_senha
export DB_NAME=meu_banco

Se não configurar as variáveis, a aplicação usará os valores padrão.
*/
