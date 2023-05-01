package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/tatsuyayamauchi/golang-echo-server-example/web/pkg/cmd/web/config"
	"github.com/tatsuyayamauchi/golang-echo-server-example/web/pkg/handler"
	"github.com/tatsuyayamauchi/golang-echo-server-example/web/pkg/middleware"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Error(err)
		panic(err)
	}
	// Create a new Echo instance
	e := echo.New()

	// Middleware
	// TODO: どうにかする
	for _, m := range middleware.RootMiddleware(c.Domain(), c.RequestTimeout()) {
		e.Use(m)
	}

	// Routes
	loginGroup := e.Group("/v1/login")
	loginGroup.POST("", handler.LoginHandlerFunc(c.JwtExpiredHour(), c.JwtSecret()))

	v1Group := e.Group("/v1")
	v1Group.Use(middleware.JwtAuth(c.JwtSecret()))
	v1Group.GET("/hello", handler.HelloHandlerFunc("v1/hello handler"))
	v1Group.GET("/hello2", handler.HelloHandlerFunc("v1/hello2 handler"))

	if c.IsEnableDebug() {
		pprofGroup := e.Group("/debug/pprof")
		for path, h := range handler.DebugHandlerFuncs() {
			pprofGroup.Any(path, h)
		}
	}

	// Start the server
	if err := e.Start(":8080"); err != nil {
		log.Error(err)
	}
}
