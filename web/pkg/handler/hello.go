package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloHandlerFunc(str string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"data": "Hello, " + str})
	}
}
