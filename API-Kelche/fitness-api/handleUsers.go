package handlers // Pacote handlers contém as funções que respondem às requisições HTTP

/*
Este arquivo (handleUsers.go) contém todas as funções que processam requisições relacionadas a usuários.
Cada função é um "handler" que recebe uma requisição HTTP e retorna uma resposta.

Conceitos importantes:
- Handler: função que processa requisições HTTP
- echo.Context: contém informações da requisição (dados, parâmetros, etc.)
- HTTP Status Codes: códigos que indicam o resultado da operação
- JSON: formato de dados usado para comunicação cliente-servidor
- CRUD: Create, Read, Update, Delete (operações básicas de banco de dados)
*/

import (
	"fitness-api/models"     // Estruturas de dados (User)
	"fitness-api/repositories" // Funções de acesso ao banco de dados
	"net/http"               // Constantes HTTP (StatusOK, StatusBadRequest, etc.)
	"strconv"                // Para converter strings em números

	"github.com/labstack/echo/v4" // Framework web Echo
)

// HandleCreateUser processa requisições POST para criar novos usuários
// Recebe dados do usuário em JSON e salva no banco de dados
func HandleCreateUser(c echo.Context) error {
	// Cria uma variável vazia do tipo User
	// Esta variável será preenchida com os dados da requisição
	user := models.User{}
	
	// c.Bind() converte o JSON da requisição para a struct User
	// Se houver erro na conversão, retorna erro 400 (Bad Request)
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados Inválidos"})
	}

	// Chama a função do repositório para salvar o usuário no banco
	// repositories.CreateUser() executa o INSERT no banco de dados
	createdUser, err := repositories.CreateUser(user)
	if err != nil {
		// Se houver erro no banco, retorna erro 500 (Internal Server Error)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Retorna o usuário criado com status 201 (Created)
	// Status 201 indica que um novo recurso foi criado com sucesso
	return c.JSON(http.StatusCreated, createdUser)
}

// HandleGetUsers processa requisições GET para listar todos os usuários
// Retorna uma lista com todos os usuários cadastrados no banco
func HandleGetUsers(c echo.Context) error {
	// Chama a função do repositório para buscar todos os usuários
	// repositories.GetUsers() executa o SELECT * FROM users
	users, err := repositories.GetUsers()
	if err != nil {
		// Se houver erro no banco, retorna erro 500
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Retorna a lista de usuários com status 200 (OK)
	return c.JSON(http.StatusOK, users)
}

// HandleGetUser processa requisições GET para buscar um usuário específico
// Recebe o ID do usuário na URL (ex: /users/123) e retorna os dados dele
func HandleGetUser(c echo.Context) error {
	// c.Param("id") extrai o parâmetro "id" da URL
	// Exemplo: se a URL for /users/123, id será "123"
	id := c.Param("id")
	
	// strconv.Atoi() converte string para int
	// Se a conversão falhar, retorna erro 400 (ID inválido)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID Inválido"})
	}

	// Busca o usuário no banco de dados pelo ID
	user, err := repositories.GetUser(idInt)
	if err != nil {
		// Se o usuário não for encontrado, retorna erro 404 (Not Found)
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Usuário Não Encontrado"})
	}

	// Retorna o usuário encontrado com status 200 (OK)
	return c.JSON(http.StatusOK, user)
}

// HandleUpdateUser processa requisições PUT para atualizar usuários existentes
// Recebe o ID na URL e os novos dados em JSON
func HandleUpdateUser(c echo.Context) error {
	// Extrai e converte o ID da URL
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID Inválido"})
	}

	// Converte o JSON da requisição para a struct User
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados Inválidos"})
	}

	// Atualiza o usuário no banco de dados
	updateUser, err := repositories.UpdateUser(user, idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Retorna o usuário atualizado com status 200 (OK)
	return c.JSON(http.StatusOK, updateUser)
}

// HandleDeleteUser processa requisições DELETE para remover usuários
// Recebe o ID na URL e remove o usuário do banco de dados
func HandleDeleteUser(c echo.Context) error {
	// Extrai e converte o ID da URL
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID Inválido"})
	}

	// Remove o usuário do banco de dados
	err = repositories.DeleteUser(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Retorna mensagem de sucesso com status 200 (OK)
	return c.JSON(http.StatusOK, map[string]string{"message": "Usuário Deletado com Sucesso"})
}

/*
CÓDIGOS DE STATUS HTTP UTILIZADOS:

- 200 (OK): Requisição bem-sucedida
- 201 (Created): Recurso criado com sucesso
- 400 (Bad Request): Dados inválidos na requisição
- 404 (Not Found): Recurso não encontrado
- 500 (Internal Server Error): Erro interno do servidor

EXEMPLO DE REQUISIÇÃO POST /users:

curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva",
    "email": "joao@email.com",
    "password": "senha123"
  }'

RESPOSTA ESPERADA:

{
  "id": 1,
  "name": "João Silva",
  "email": "joao@email.com",
  "password": "senha123",
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
*/