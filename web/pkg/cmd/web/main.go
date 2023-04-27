package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/tatsuyayamauchi/golang-echo-server-example/web/pkg/handler"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", handler.HelloHandler)

	// Start the server
	e.Start(":8080")
}
