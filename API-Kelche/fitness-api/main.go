package main

import (
	storage "fitness-api"
	"fitness-api/cmd/handlers" // E o handlers serve para criar as rotas da API.
	"fitness-api/cmd/storage"

	"github.com/labstack/echo/v4"            // Echo serve para criar um servidor web HTTP.
	"github.com/labstack/echo/v4/middleware" // E o middleware serve para permitir o compartilhamento de recursos entre essas diferentes origens
)

func main() {

	storage.ConnectDB()

	e := echo.New() // Instancia echo serve para criar um servidor web HTTP.


	//Middleware
	e.Use(middleware.Logger()) // Middleware para logar as requisicoes e respostas
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	})) // Middleware para permitir acesso de origens diferentes pelo CORS ( Cross-Origin Resource Sharing ) que serve para permitir o compartilhamento de recursos entre essas diferentes origens.

	// Rotas
	e.GET("/", handlers.Home)

	// Rotas de usuários
	e.GET("/users", handlers.HandleGetUsers)
	e.GET("/users/:id", handlers.HandleGetUser)
	e.POST("/users", handlers.HandleCreateUser)
	e.PUT("/users/:id", handlers.HandleUpdateUser)
	e.DELETE("/users/:id", handlers.HandleDeleteUser)

	// Rotas de Medições
	e.GET("/measurements", handlers.HandleGetMeasurements)
	e.GET("/measurements/:id", handlers.HandleGetMeasurement)
	e.POST("/measurements", handlers.HandleCreateMeasurement)
	e.PUT("/measurements/:id", handlers.HandleUpdateMeasurement)
	e.DELETE("/measurements/:id", handlers.HandleDeleteMeasurement)

	e.Logger.Fatal(e.Start(":8080")) // Inicia o servidor na porta 8080
}