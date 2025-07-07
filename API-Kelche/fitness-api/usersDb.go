package repositories // Pacote repositories contém as funções de acesso ao banco de dados

/*
Este arquivo (usersDb.go) contém todas as funções que interagem com o banco de dados para usuários.
Aqui implementamos as operações CRUD (Create, Read, Update, Delete) para a tabela users.

Conceitos importantes:
- Repository Pattern: separa a lógica de acesso a dados da lógica de negócio
- SQL: linguagem para consultar e manipular bancos de dados relacionais
- Prepared Statements: queries parametrizadas para evitar SQL injection
- Error Handling: tratamento adequado de erros do banco de dados
- Transactions: operações que devem ser executadas como uma unidade
*/

import (
	"fitness-api/models" // Estruturas de dados (User)
	"fitness-api/storage" // Conexão com o banco de dados
	"time"                // Para trabalhar com datas e horários
)

// CreateUser insere um novo usuário no banco de dados
// Recebe um User e retorna o User criado com ID preenchido
func CreateUser(user models.User) (models.User, error) {
	// Obtém a conexão com o banco de dados
	db := storage.GetDB()
	
	// Query SQL para inserir um novo usuário
	// $1, $2, $3, etc. são placeholders que serão substituídos pelos valores
	// RETURNING id retorna o ID gerado automaticamente pelo banco
	sqlStatement := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	// Executa a query com os parâmetros
	// db.QueryRow() executa uma query que retorna uma única linha
	// .Scan() copia os valores retornados para a variável user.Id
	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password, time.Now(), time.Now()).Scan(&user.Id)
	if err != nil {
		// Se houver erro, retorna um User vazio e o erro
		return models.User{}, err
	}

	// Define os timestamps de criação e atualização
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	
	// Retorna o usuário criado com ID preenchido
	return user, nil
}

// GetUsers busca todos os usuários do banco de dados
// Retorna uma lista (slice) de usuários
func GetUsers() ([]models.User, error) {
	// Obtém a conexão com o banco de dados
	db := storage.GetDB()
	
	// Query SQL para buscar todos os usuários
	sqlStatement := `SELECT id, name, email, password, created_at, updated_at FROM users`

	// Executa a query
	// db.Query() executa uma query que pode retornar múltiplas linhas
	rows, err := db.Query(sqlStatement)
	if err != nil {
		// Se houver erro, retorna nil e o erro
		return nil, err
	}
	// defer rows.Close() garante que as linhas sejam fechadas após o uso
	defer rows.Close()

	// Cria um slice vazio para armazenar os usuários
	var users []models.User 
	
	// Itera sobre cada linha retornada pela query
	for rows.Next() {
		// Cria uma variável para armazenar os dados de cada usuário
		var user models.User
		
		// rows.Scan() copia os valores da linha atual para a struct user
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		
		// Adiciona o usuário ao slice
		users = append(users, user)
	}

	// Retorna a lista de usuários
	return users, nil
}

// GetUser busca um usuário específico pelo ID
// Recebe o ID e retorna o usuário encontrado
func GetUser(id int) (models.User, error) {
	// Obtém a conexão com o banco de dados
	db := storage.GetDB()
	
	// Query SQL para buscar um usuário pelo ID
	// WHERE id = $1 filtra apenas o usuário com o ID especificado
	sqlStatement := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1`

	// Cria uma variável para armazenar os dados do usuário
	var user models.User
	
	// Executa a query e copia os resultados para a struct user
	err := db.QueryRow(sqlStatement, id).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		// Se houver erro (usuário não encontrado), retorna User vazio e o erro
		return models.User{}, err
	}

	// Retorna o usuário encontrado
	return user, nil
}

// UpdateUser atualiza os dados de um usuário existente
// Recebe os novos dados do usuário e o ID, retorna o usuário atualizado
func UpdateUser(user models.User, id int) (models.User, error) {
	// Obtém a conexão com o banco de dados
	db := storage.GetDB()
	
	// Query SQL para atualizar um usuário
	// SET define os campos que serão atualizados
	// WHERE id = $5 garante que apenas o usuário correto seja atualizado
	// RETURNING id retorna o ID do usuário atualizado
	sqlStatement := `UPDATE users SET name = $1, email = $2, password = $3, updated_at = $4 WHERE id = $5 RETURNING id`

	// Executa a query de atualização
	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password, time.Now(), id).Scan(&user.Id)
	if err != nil {
		// Se houver erro, retorna User vazio e o erro
		return models.User{}, err
	}

	// Define o timestamp de atualização
	user.UpdatedAt = time.Now()
	
	// Retorna o usuário atualizado
	return user, nil
}

// DeleteUser remove um usuário do banco de dados
// Recebe o ID do usuário e remove a linha correspondente
func DeleteUser(id int) error {
	// Obtém a conexão com o banco de dados
	db := storage.GetDB()
	
	// Query SQL para deletar um usuário
	// WHERE id = $1 garante que apenas o usuário correto seja removido
	sqlStatement := `DELETE FROM users WHERE id = $1`

	// Executa a query de deleção
	// db.Exec() executa uma query que não retorna linhas
	_, err := db.Exec(sqlStatement, id)
	
	// Retorna o erro (se houver)
	return err
}

/*
EXEMPLOS DE QUERIES SQL UTILIZADAS:

1. INSERT (CreateUser):
   INSERT INTO users (name, email, password, created_at, updated_at) 
   VALUES ($1, $2, $3, $4, $5) RETURNING id

2. SELECT (GetUsers):
   SELECT id, name, email, password, created_at, updated_at FROM users

3. SELECT com WHERE (GetUser):
   SELECT id, name, email, password, created_at, updated_at 
   FROM users WHERE id = $1

4. UPDATE (UpdateUser):
   UPDATE users SET name = $1, email = $2, password = $3, updated_at = $4 
   WHERE id = $5 RETURNING id

5. DELETE (DeleteUser):
   DELETE FROM users WHERE id = $1

CONCEITOS DE SEGURANÇA:

- Prepared Statements ($1, $2, etc.): evitam SQL injection
- Validação de entrada: sempre validar dados antes de inserir no banco
- Controle de acesso: implementar autenticação e autorização
- Criptografia: senhas devem ser hasheadas (bcrypt, scrypt, etc.)

MELHORIAS SUGERIDAS:

1. Implementar soft delete (não deletar, apenas marcar como inativo)
2. Adicionar logs de auditoria (quem alterou, quando, etc.)
3. Implementar cache para consultas frequentes
4. Adicionar índices no banco para melhor performance
5. Implementar paginação para listas grandes
*/