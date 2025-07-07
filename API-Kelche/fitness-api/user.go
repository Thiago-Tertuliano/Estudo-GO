package models // Pacote models contém as estruturas de dados (structs) da aplicação

/*
Este arquivo (user.go) define a estrutura de dados para usuários da aplicação.
Em Go, usamos structs para representar objetos/entidades do nosso sistema.

Conceitos importantes:
- Struct: é como uma "classe" em outras linguagens, define a estrutura de um objeto
- Tags JSON: definem como os campos serão serializados/deserializados em JSON
- Package models: agrupa todas as estruturas de dados da aplicação
*/

import (
	"time" // Biblioteca para trabalhar com datas e horários
)

// User representa um usuário no sistema
// Esta struct define todos os campos que um usuário pode ter
type User struct {
	// ID único do usuário (chave primária no banco de dados)
	// int = número inteiro
	// `json:"id"` = quando converter para JSON, este campo será chamado "id"
	Id int `json:"id"`

	// Nome completo do usuário
	// string = texto
	// `json:"name"` = no JSON será "name"
	Name string `json:"name"`

	// Email do usuário (deve ser único)
	// Usado para login e identificação
	Email string `json:"email"`

	// Senha do usuário (em produção deve ser criptografada!)
	// ⚠️ IMPORTANTE: Nunca retorne a senha em respostas da API
	Password string `json:"password"`

	// Data e hora de criação do usuário
	// time.Time = tipo específico para datas em Go
	// `json:"created_at"` = no JSON será "created_at"
	CreatedAt time.Time `json:"created_at"`

	// Data e hora da última atualização do usuário
	// Útil para saber quando o usuário foi modificado pela última vez
	UpdatedAt time.Time `json:"updated_at"`
}

/*
EXEMPLO DE JSON DE UM USUÁRIO:

{
  "id": 1,
  "name": "João Silva",
  "email": "joao@email.com",
  "password": "senha123",
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}

COMO FUNCIONAM AS TAGS JSON:

- `json:"nome_campo"` = define o nome do campo no JSON
- `json:"-"` = exclui o campo do JSON (útil para senhas)
- `json:"nome,omitempty"` = omite o campo se estiver vazio
- `json:",omitempty"` = mantém o nome original mas omite se vazio

EXEMPLO DE USO:

user := User{
    Name: "Maria",
    Email: "maria@email.com",
    Password: "123456",
}

// Quando convertido para JSON:
// {"name":"Maria","email":"maria@email.com","password":"123456"}
*/