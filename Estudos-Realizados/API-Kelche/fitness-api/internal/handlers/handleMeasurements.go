package handlers // Pacote handlers contém as funções que respondem às requisições HTTP

/*
Este arquivo (handleMeasurements.go) contém todas as funções que processam requisições relacionadas a medições de fitness.
Cada função é um "handler" que recebe uma requisição HTTP e retorna uma resposta.

Conceitos importantes:
- Handler: função que processa requisições HTTP
- echo.Context: contém informações da requisição (dados, parâmetros, etc.)
- HTTP Status Codes: códigos que indicam o resultado da operação
- JSON: formato de dados usado para comunicação cliente-servidor
- CRUD: Create, Read, Update, Delete (operações básicas de banco de dados)
- Relacionamentos: medições pertencem a usuários (user_id)
*/

import (
	"fitness-api/internal/models"     // Estruturas de dados (Measurements)
	"fitness-api/internal/repositories" // Funções de acesso ao banco de dados
	"net/http"               // Constantes HTTP (StatusOK, StatusBadRequest, etc.)
	"strconv"                // Para converter strings em números

	"github.com/labstack/echo/v4" // Framework web Echo
)

// HandleCreateMeasurement processa requisições POST para criar novas medições
// Recebe dados da medição em JSON e salva no banco de dados
func HandleCreateMeasurement(c echo.Context) error {
	// Cria uma variável vazia do tipo Measurements
	// Esta variável será preenchida com os dados da requisição
	measurement := models.Measurements{}
	
	// c.Bind() converte o JSON da requisição para a struct Measurements
	// Se houver erro na conversão, retorna erro 400 (Bad Request)
	if err := c.Bind(&measurement); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados Inválidos"})
	}

	// Chama a função do repositório para salvar a medição no banco
	// repositories.CreateMeasurement() executa o INSERT no banco de dados
	createdMeasurement, err := repositories.CreateMeasurement(measurement)
	if err != nil {
		// Se houver erro no banco, retorna erro 500 (Internal Server Error)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Retorna a medição criada com status 201 (Created)
	// Status 201 indica que um novo recurso foi criado com sucesso
	return c.JSON(http.StatusCreated, createdMeasurement)
}

// HandleGetMeasurements processa requisições GET para listar todas as medições
// Retorna uma lista com todas as medições cadastradas no banco
func HandleGetMeasurements(c echo.Context) error {
	// Chama a função do repositório para buscar todas as medições
	// repositories.GetMeasurements() executa o SELECT * FROM measurements
	measurements, err := repositories.GetMeasurements()
	if err != nil {
		// Se houver erro no banco, retorna erro 500
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Retorna a lista de medições com status 200 (OK)
	return c.JSON(http.StatusOK, measurements)
}

// HandleGetMeasurement processa requisições GET para buscar uma medição específica
// Recebe o ID da medição na URL (ex: /measurements/123) e retorna os dados dela
func HandleGetMeasurement(c echo.Context) error {
	// c.Param("id") extrai o parâmetro "id" da URL
	// Exemplo: se a URL for /measurements/123, id será "123"
	id := c.Param("id")
	
	// strconv.Atoi() converte string para int
	// Se a conversão falhar, retorna erro 400 (ID inválido)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID Inválido"})
	}

	// Busca a medição no banco de dados pelo ID
	measurement, err := repositories.GetMeasurement(idInt)
	if err != nil {
		// Se a medição não for encontrada, retorna erro 404 (Not Found)
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Medida Não Encontrada"})
	}

	// Retorna a medição encontrada com status 200 (OK)
	return c.JSON(http.StatusOK, measurement)
}

// HandleUpdateMeasurement processa requisições PUT para atualizar medições existentes
// Recebe o ID na URL e os novos dados em JSON
func HandleUpdateMeasurement(c echo.Context) error {
	// Extrai e converte o ID da URL
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID Inválido"})
	}

	// Converte o JSON da requisição para a struct Measurements
	measurement := models.Measurements{}
	if err := c.Bind(&measurement); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados Inválidos"})
	}

	// Atualiza a medição no banco de dados
	updateMeasurement, err := repositories.UpdateMeasurement(measurement, idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Retorna a medição atualizada com status 200 (OK)
	return c.JSON(http.StatusOK, updateMeasurement)
}

// HandleDeleteMeasurement processa requisições DELETE para remover medições
// Recebe o ID na URL e remove a medição do banco de dados
func HandleDeleteMeasurement(c echo.Context) error {
	// Extrai e converte o ID da URL
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID Inválido"})
	}

	// Remove a medição do banco de dados
	err = repositories.DeleteMeasurement(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Retorna mensagem de sucesso com status 200 (OK)
	return c.JSON(http.StatusOK, map[string]string{"message": "Medida Deletada com Sucesso"})
}

/*
CÓDIGOS DE STATUS HTTP UTILIZADOS:

- 200 (OK): Requisição bem-sucedida
- 201 (Created): Recurso criado com sucesso
- 400 (Bad Request): Dados inválidos na requisição
- 404 (Not Found): Recurso não encontrado
- 500 (Internal Server Error): Erro interno do servidor

EXEMPLO DE REQUISIÇÃO POST /measurements:

curl -X POST http://localhost:8080/measurements \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "weight": 75.5,
    "height": 1.75,
    "body_fat": 15.2
  }'

RESPOSTA ESPERADA:

{
  "id": 1,
  "user_id": 1,
  "weight": 75.5,
  "height": 1.75,
  "body_fat": 15.2,
  "created_at": "2024-01-15T14:30:00Z"
}

RELACIONAMENTO COM USUÁRIOS:

- user_id deve referenciar um usuário existente na tabela users
- Se o user_id não existir, a inserção falhará (restrição de chave estrangeira)
- Uma medição sempre pertence a um usuário específico

VALIDAÇÕES SUGERIDAS (para implementar):

1. Verificar se o user_id existe antes de criar a medição
2. Validar se weight, height e body_fat estão em ranges realistas
3. Verificar se não há medição duplicada para o mesmo usuário na mesma data
4. Implementar autenticação para garantir que usuários só vejam suas próprias medições
*/