package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOk, "API de Fitness - GO + Echo + PostgreeSQL")
}
