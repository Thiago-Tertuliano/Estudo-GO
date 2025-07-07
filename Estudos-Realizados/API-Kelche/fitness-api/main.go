package main

/*
Este é o arquivo principal (main.go) da nossa API de Fitness.
Em Go, o arquivo main.go é o ponto de entrada da aplicação - é onde tudo começa!

O que este arquivo faz:
1. Configura o servidor web usando o framework Echo
2. Estabelece conexão com o banco de dados
3. Configura middlewares (filtros que processam requisições)
4. Define todas as rotas da API
5. Inicia o servidor na porta 8080
*/

import (
	"fitness-api/internal/handlers" // Pacote que contém todas as funções que respondem às requisições HTTP
	"fitness-api/db"  // Pacote que gerencia a conexão com o banco de dados

	"github.com/labstack/echo/v4"            // Framework web Echo - facilita a criação de APIs REST
	"github.com/labstack/echo/v4/middleware" // Middlewares do Echo - funcionalidades extras para o servidor
	"fmt"
)

// Função main() é obrigatória em Go - é onde a execução do programa começa
func main() {
	// Estabelece conexão com o banco de dados PostgreSQL
	// Esta função está definida no arquivo database.go
	storage.ConnectDB()

	// Cria uma nova instância do servidor Echo
	// Echo é um framework web que facilita a criação de APIs REST
	e := echo.New()

	// ===== CONFIGURAÇÃO DE MIDDLEWARES =====
	// Middlewares são funções que executam antes ou depois das requisições
	// Eles podem modificar requisições, adicionar logs, verificar autenticação, etc.

	// Middleware de Logger - registra todas as requisições e respostas no console
	// Útil para debug e monitoramento da API
	e.Use(middleware.Logger())

	// Middleware CORS - permite que aplicações web de outros domínios acessem nossa API
	// CORS = Cross-Origin Resource Sharing
	// AllowOrigins: []string{"*"} = permite acesso de qualquer origem (não recomendado para produção)
	// AllowHeaders = define quais cabeçalhos HTTP são permitidos
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // ⚠️ Em produção, especifique domínios específicos
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// ===== DEFINIÇÃO DAS ROTAS =====
	// Cada rota mapeia um caminho URL para uma função handler
	// GET = buscar dados, POST = criar, PUT = atualizar, DELETE = remover

	// Rota raiz - página inicial da API
	e.GET("/", handlers.Home)

	// ===== ROTAS DE USUÁRIOS =====
	// CRUD completo para usuários (Create, Read, Update, Delete)
	
	// GET /users - Lista todos os usuários
	e.GET("/users", handlers.HandleGetUsers)
	
	// GET /users/:id - Busca um usuário específico pelo ID
	// :id é um parâmetro dinâmico que será substituído pelo ID real
	e.GET("/users/:id", handlers.HandleGetUser)
	
	// POST /users - Cria um novo usuário
	e.POST("/users", handlers.HandleCreateUser)
	
	// PUT /users/:id - Atualiza um usuário existente
	e.PUT("/users/:id", handlers.HandleUpdateUser)
	
	// DELETE /users/:id - Remove um usuário
	e.DELETE("/users/:id", handlers.HandleDeleteUser)

	// ===== ROTAS DE MEDIÇÕES =====
	// CRUD completo para medições de fitness (peso, altura, gordura corporal)
	
	// GET /measurements - Lista todas as medições
	e.GET("/measurements", handlers.HandleGetMeasurements)
	
	// GET /measurements/:id - Busca uma medição específica pelo ID
	e.GET("/measurements/:id", handlers.HandleGetMeasurement)
	
	// POST /measurements - Cria uma nova medição
	e.POST("/measurements", handlers.HandleCreateMeasurement)
	
	// PUT /measurements/:id - Atualiza uma medição existente
	e.PUT("/measurements/:id", handlers.HandleUpdateMeasurement)
	
	// DELETE /measurements/:id - Remove uma medição
	e.DELETE("/measurements/:id", handlers.HandleDeleteMeasurement)

	// ===== INICIALIZAÇÃO DO SERVIDOR =====
	// Inicia o servidor na porta 8080
	// e.Start() inicia o servidor HTTP
	// e.Logger.Fatal() para a aplicação se houver erro na inicialização
	fmt.Println("🚀 Servidor iniciado em http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}