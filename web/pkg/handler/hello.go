package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type resp struct {
	Data string `json:"data"`
}

func HelloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, &resp{Data: "This is HelloHandler"})
}
