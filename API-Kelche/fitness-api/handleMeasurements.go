package handlers

import (
	"fitnes-api/cmd/repositories"
	"fitness-api/cmd/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func HandleCreateMeasurement(c echo.Context) error {
	measurement := models.Measurements{}
	if err := c.Bind(&measurement); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados Inválidos"})
	}

	createdMeasurement, err := repositories.CreateMeasurement(measurement)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdMeasurement)
}

func HandleGetMeasurements(c echo.Context) error {
	measurements, err := repositories.GetMeasurements()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, measurements)
}

func HandleGetMeasurement(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID Inválido"})
	}

	measurement, err := repositories.GetMeasurement(idInt)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Medida Não Encontrada"})
	}

	return c.JSON(http.StatusOK, measurement)
}

func HandleUpdateMeasurement(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID Inválido"})
	}

	measurement := models.Measurements{}
	if err := c.Bind(&measurement); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados Inválidos"})
	}

	updateMeasurement, err := repositories.UpdateMeasurement(measurement, idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updateMeasurement)
}

func HandleDeleteMeasurement(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID Inválido"})
	}

	err = repositories.DeleteMeasurement(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Medida Deletada com Sucesso"})
}