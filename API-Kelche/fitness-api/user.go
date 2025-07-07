package models // Pacote models do arquivo user.go

import (
	"time" // Biblioteca para trabalhar com datas
)

type User struct {
	Id int `json:"id"` // ID do usuário
	Name string `json"name"` // Nome do usuário
	Email string `json"email"` // E-mail do usuário
	Password string `json"password"` // Senha do usuário
	CreatedAt time.Time `json"created_at"` // Criação do usuário
	UptadeAt time.Time `json"uptade_at"` // Atualização do usuário

}