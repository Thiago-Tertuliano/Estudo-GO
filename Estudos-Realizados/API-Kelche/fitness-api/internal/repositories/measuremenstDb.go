package repositories // Pacote repositories contém as funções de acesso ao banco de dados

/*
Este arquivo (measuremenstDb.go) contém todas as funções que interagem com o banco de dados para medições.
Aqui implementamos as operações CRUD (Create, Read, Update, Delete) para a tabela measurements.

Conceitos importantes:
- Repository Pattern: separa a lógica de acesso a dados da lógica de negócio
- Relacionamentos: medições pertencem a usuários (chave estrangeira user_id)
- SQL: linguagem para consultar e manipular bancos de dados relacionais
- Prepared Statements: queries parametrizadas para evitar SQL injection
- Error Handling: tratamento adequado de erros do banco de dados
*/

import (
	"fitness-api/internal/models" // Estruturas de dados (Measurements)
	"fitness-api/db" // Conexão com o banco de dados
	"time"                // Para trabalhar com datas e horários
)

// CreateMeasurement insere uma nova medição no banco de dados
// Recebe um Measurements e retorna o Measurements criado com ID preenchido
func CreateMeasurement(measurement models.Measurements) (models.Measurements, error) {
	// Obtém a conexão com o banco de dados
	db := storage.GetDB()
	
	// Query SQL para inserir uma nova medição
	// $1, $2, $3, etc. são placeholders que serão substituídos pelos valores
	// RETURNING id retorna o ID gerado automaticamente pelo banco
	sqlStatement := `INSERT INTO measurements (user_id, weight, height, body_fat, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	// Executa a query com os parâmetros
	// db.QueryRow() executa uma query que retorna uma única linha
	// .Scan() copia os valores retornados para a variável measurement.Id
	err := db.QueryRow(sqlStatement, measurement.UserId, measurement.Weight, measurement.Height, measurement.BodyFat, time.Now()).Scan(&measurement.Id)
	if err != nil {
		// Se houver erro, retorna um Measurements vazio e o erro
		return models.Measurements{}, err
	}

	// Define o timestamp de criação
	measurement.CreatedAt = time.Now()
	
	// Retorna a medição criada com ID preenchido
	return measurement, nil
}

// GetMeasurements busca todas as medições do banco de dados
// Retorna uma lista (slice) de medições
func GetMeasurements() ([]models.Measurements, error) {
	// Obtém a conexão com o banco de dados
	db := storage.GetDB()
	
	// Query SQL para buscar todas as medições
	sqlStatement := `SELECT id, user_id, weight, height, body_fat, created_at FROM measurements`

	// Executa a query
	// db.Query() executa uma query que pode retornar múltiplas linhas
	rows, err := db.Query(sqlStatement)
	if err != nil {
		// Se houver erro, retorna nil e o erro
		return nil, err
	}
	// defer rows.Close() garante que as linhas sejam fechadas após o uso
	defer rows.Close()

	// Cria um slice vazio para armazenar as medições
	var measurements []models.Measurements
	
	// Itera sobre cada linha retornada pela query
	for rows.Next() {
		// Cria uma variável para armazenar os dados de cada medição
		var measurement models.Measurements
		
		// rows.Scan() copia os valores da linha atual para a struct measurement
		err := rows.Scan(&measurement.Id, &measurement.UserId, &measurement.Weight, &measurement.Height, &measurement.BodyFat, &measurement.CreatedAt)
		if err != nil {
			return nil, err
		}
		
		// Adiciona a medição ao slice
		measurements = append(measurements, measurement)
	}

	// Retorna a lista de medições
	return measurements, nil
}

// GetMeasurement busca uma medição específica pelo ID
// Recebe o ID e retorna a medição encontrada
func GetMeasurement(id int) (models.Measurements, error) {
	// Obtém a conexão com o banco de dados
	db := storage.GetDB()
	
	// Query SQL para buscar uma medição pelo ID
	// WHERE id = $1 filtra apenas a medição com o ID especificado
	sqlStatement := `SELECT id, user_id, weight, height, body_fat, created_at FROM measurements WHERE id = $1`

	// Cria uma variável para armazenar os dados da medição
	var measurement models.Measurements
	
	// Executa a query e copia os resultados para a struct measurement
	err := db.QueryRow(sqlStatement, id).Scan(&measurement.Id, &measurement.UserId, &measurement.Weight, &measurement.Height, &measurement.BodyFat, &measurement.CreatedAt)
	if err != nil {
		// Se houver erro (medição não encontrada), retorna Measurements vazio e o erro
		return models.Measurements{}, err
	}

	// Retorna a medição encontrada
	return measurement, nil
}

// UpdateMeasurement atualiza os dados de uma medição existente
// Recebe os novos dados da medição e o ID, retorna a medição atualizada
func UpdateMeasurement(measurement models.Measurements, id int) (models.Measurements, error) {
	// Obtém a conexão com o banco de dados
	db := storage.GetDB()
	
	// Query SQL para atualizar uma medição
	// SET define os campos que serão atualizados
	// WHERE id = $5 garante que apenas a medição correta seja atualizada
	// RETURNING id retorna o ID da medição atualizada
	sqlStatement := `UPDATE measurements SET user_id = $1, weight = $2, height = $3, body_fat = $4 WHERE id = $5 RETURNING id`

	// Executa a query de atualização
	err := db.QueryRow(sqlStatement, measurement.UserId, measurement.Weight, measurement.Height, measurement.BodyFat, id).Scan(&measurement.Id)
	if err != nil {
		// Se houver erro, retorna Measurements vazio e o erro
		return models.Measurements{}, err
	}

	// Retorna a medição atualizada
	return measurement, nil
}

// DeleteMeasurement remove uma medição do banco de dados
// Recebe o ID da medição e remove a linha correspondente
func DeleteMeasurement(id int) error {
	// Obtém a conexão com o banco de dados
	db := storage.GetDB()
	
	// Query SQL para deletar uma medição
	// WHERE id = $1 garante que apenas a medição correta seja removida
	sqlStatement := `DELETE FROM measurements WHERE id = $1`

	// Executa a query de deleção
	// db.Exec() executa uma query que não retorna linhas
	_, err := db.Exec(sqlStatement, id)
	
	// Retorna o erro (se houver)
	return err
}

/*
EXEMPLOS DE QUERIES SQL UTILIZADAS:

1. INSERT (CreateMeasurement):
   INSERT INTO measurements (user_id, weight, height, body_fat, created_at) 
   VALUES ($1, $2, $3, $4, $5) RETURNING id

2. SELECT (GetMeasurements):
   SELECT id, user_id, weight, height, body_fat, created_at FROM measurements

3. SELECT com WHERE (GetMeasurement):
   SELECT id, user_id, weight, height, body_fat, created_at 
   FROM measurements WHERE id = $1

4. UPDATE (UpdateMeasurement):
   UPDATE measurements SET user_id = $1, weight = $2, height = $3, body_fat = $4 
   WHERE id = $5 RETURNING id

5. DELETE (DeleteMeasurement):
   DELETE FROM measurements WHERE id = $1

RELACIONAMENTO COM USUÁRIOS:

- user_id é uma chave estrangeira que referencia users(id)
- Se um usuário for deletado, suas medições também serão (CASCADE)
- Uma medição sempre pertence a um usuário válido

QUERIES ÚTEIS ADICIONAIS (para implementar):

1. Buscar medições de um usuário específico:
   SELECT * FROM measurements WHERE user_id = $1 ORDER BY created_at DESC

2. Buscar última medição de um usuário:
   SELECT * FROM measurements WHERE user_id = $1 ORDER BY created_at DESC LIMIT 1

3. Calcular evolução do peso de um usuário:
   SELECT weight, created_at FROM measurements 
   WHERE user_id = $1 ORDER BY created_at ASC

4. Buscar usuários com mais medições:
   SELECT user_id, COUNT(*) as total_measurements 
   FROM measurements GROUP BY user_id ORDER BY total_measurements DESC

VALIDAÇÕES IMPORTANTES:

1. Verificar se user_id existe antes de criar medição
2. Validar ranges realistas para weight, height e body_fat
3. Evitar medições duplicadas no mesmo dia para o mesmo usuário
4. Implementar autenticação para garantir privacidade dos dados

MELHORIAS SUGERIDAS:

1. Adicionar campo updated_at na tabela measurements
2. Implementar soft delete para medições
3. Criar índices para user_id e created_at
4. Implementar cache para consultas frequentes
5. Adicionar paginação para listas grandes
6. Implementar busca por período de datas
*/