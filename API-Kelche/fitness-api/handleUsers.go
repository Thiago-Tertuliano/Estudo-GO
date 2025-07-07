package handlers

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/repositories"
	"net/http"
	"os/user"
	"strconv"

	"github.com/labstack/echo/v4"
)

func HandleCreatedUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados Inválidos"})
	}

	createUser, err := repositories.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdUser)
}

func HandleGetUsers(c echo.Context) error {
	users, err := repositories.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Erro()})
	}

	return c.JSON(http.StatusOK, users)
}

func HandleGetUser(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID Inválido"})
	}

	user, err := repositories.GetUser(idInt)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Usuário Não Encontrado"})
	}

	return c.JSON(http.StatusOK, user)
}

func HandleUpdateUser(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID Inválido"})
	}

	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados Inválidos"})
	}

	updateUser, err := repositories.UpdateUser(user, idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updateUser)
}

func HandleDeleteUser(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID Inválido"})
	}

	err = repositories.DeleteUser(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Usuário Deletado com Sucesso"})
}