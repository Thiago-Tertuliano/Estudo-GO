package handlers // Pacote handlers contém as funções que respondem às requisições HTTP

/*
Este arquivo (rootHandlers.go) contém o handler para a rota raiz da API.
A rota raiz (/) é a primeira página que os usuários veem quando acessam a API.

Conceitos importantes:
- Rota raiz: endpoint principal da API (ex: http://localhost:8080/)
- Handler simples: função que retorna uma resposta básica
- Status codes: códigos HTTP que indicam o resultado da operação
- String response: resposta em formato texto simples
*/

import (
	"net/http" // Constantes HTTP (StatusOK, etc.)
	"github.com/labstack/echo/v4" // Framework web Echo
)

// Home é o handler para a rota raiz da API
// Esta função é chamada quando alguém acessa http://localhost:8080/
func Home(c echo.Context) error {
	// c.String() retorna uma resposta em formato texto
	// http.StatusOK = código 200 (sucesso)
	// A mensagem é exibida no navegador ou cliente HTTP
	return c.String(http.StatusOK, "API de Fitness - GO + Echo + PostgreSQL")
}

/*
EXEMPLOS DE RESPOSTAS:

1. Resposta atual:
   "API de Fitness - GO + Echo + PostgreSQL"

2. Resposta mais elaborada (sugestão):
   "🚀 API de Fitness v1.0
    Tecnologias: Go + Echo + PostgreSQL
    Status: Online
    Endpoints disponíveis:
    - GET /users - Listar usuários
    - POST /users - Criar usuário
    - GET /measurements - Listar medições
    - POST /measurements - Criar medição"

3. Resposta em JSON (sugestão):
   {
     "name": "API de Fitness",
     "version": "1.0.0",
     "status": "online",
     "technologies": ["Go", "Echo", "PostgreSQL"],
     "endpoints": {
       "users": "/users",
       "measurements": "/measurements"
     }
   }

CÓDIGOS DE STATUS HTTP:

- 200 (OK): Requisição bem-sucedida
- 404 (Not Found): Página não encontrada
- 500 (Internal Server Error): Erro interno do servidor

TESTANDO A ROTA:

1. Via navegador:
   Acesse: http://localhost:8080/

2. Via curl:
   curl http://localhost:8080/

3. Via Postman:
   GET http://localhost:8080/

MELHORIAS SUGERIDAS:

1. Adicionar informações da versão da API
2. Incluir links para documentação
3. Mostrar status do banco de dados
4. Adicionar timestamp da resposta
5. Implementar health check
6. Criar página de documentação interativa (Swagger)
*/
