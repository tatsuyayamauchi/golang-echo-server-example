package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	echo_mw "github.com/labstack/echo/v4/middleware"

	"github.com/tatsuyayamauchi/golang-echo-server-example/web/pkg/cmd/web/config"
	"github.com/tatsuyayamauchi/golang-echo-server-example/web/pkg/middleware"
)

type Router struct {
	c *config.Config
	e *echo.Echo
}

func NewRouter(c *config.Config) *Router {
	e := echo.New()

	// register middlewares
	e.Use(echo_mw.Logger())
	e.Use(echo_mw.Recover())
	e.Use(middleware.Timeout(c.RequestTimeout()))
	e.Use(middleware.Cors(c.Domain()))
	e.Use(middleware.Gzip(c.CompressionLevel()))

	// register route
	loginGroup := e.Group("/v1/login")
	loginGroup.POST("", LoginHandlerFunc(c.JwtExpiredHour(), c.JwtSecret()))

	v1Group := e.Group("/v1")
	v1Group.Use(middleware.JwtAuth(c.JwtSecret()))
	v1Group.GET("/hello", HelloHandlerFunc("v1/hello handler"))
	v1Group.GET("/hello2", HelloHandlerFunc("v1/hello2 handler"))

	if c.IsEnableDebug() {
		pprofGroup := e.Group("/debug/pprof")
		for path, h := range DebugHandlerFuncs() {
			pprofGroup.Any(path, h)
		}
	}

	return &Router{
		c: c,
		e: e,
	}
}

func (r *Router) StartHTTPServer() error {
	return r.e.Start(r.portAddr())
}

// see: https://echo.labstack.com/guide/http_server/#http-server
func (r *Router) portAddr() string {
	return fmt.Sprintf(":%s", r.c.Port())
}
