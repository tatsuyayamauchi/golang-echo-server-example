package main

import (
	"net/http"
	"net/http/pprof"

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

	// pprof
	pprofGroup := e.Group("/debug/pprof")
	pprofGroup.Any("/cmdline", echo.WrapHandler(http.HandlerFunc(pprof.Cmdline)))
	pprofGroup.Any("/profile", echo.WrapHandler(http.HandlerFunc(pprof.Profile)))
	pprofGroup.Any("/symbol", echo.WrapHandler(http.HandlerFunc(pprof.Symbol)))
	pprofGroup.Any("/trace", echo.WrapHandler(http.HandlerFunc(pprof.Trace)))
	pprofGroup.Any("/*", echo.WrapHandler(http.HandlerFunc(pprof.Index)))

	// Start the server
	e.Start(":8080")
}
