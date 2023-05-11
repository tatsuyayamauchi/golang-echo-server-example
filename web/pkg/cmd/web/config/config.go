package config

import (
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	echo_mw "github.com/labstack/echo/v4/middleware"

	"github.com/tatsuyayamauchi/golang-echo-server-example/web/pkg/handler"
	"github.com/tatsuyayamauchi/golang-echo-server-example/web/pkg/middleware"
)

const (
	compressionLevel int   = 5
	jwtExpiredHour   int32 = 72
)

type Config struct {
	domain      string
	port        string
	enableDebug bool
	reqTimeout  time.Duration
	jwtSecret   string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func NewConfig() (*Config, error) {
	domain, ok := os.LookupEnv("DOMAIN")
	if !ok {
		return nil, fmt.Errorf("DOMAIN environment variable is not set")
	}
	port := getEnv("PORT", "8080")

	jwtSecret, ok := os.LookupEnv("JWT_SECRET_KEY")
	if !ok {
		return nil, fmt.Errorf("JWT_SECRET_KEY environment variable is not set")
	}

	reqTimeout, err := time.ParseDuration(getEnv("REQUEST_TIMEOUT_SEC", "10") + "s")
	if err != nil {
		return nil, fmt.Errorf("parse REQUEST_TIMEOUT_SEC error, err: %v", err)
	}
	enableDebug := getEnv("ENABLE_DEBUG", "false") == "true"

	return &Config{
		domain:      domain,
		port:        port,
		enableDebug: enableDebug,
		reqTimeout:  reqTimeout,
		jwtSecret:   jwtSecret,
	}, nil
}

// 環境変数受け取り
func (c *Config) Domain() string                { return c.domain }
func (c *Config) Port() string                  { return c.port }
func (c *Config) IsEnableDebug() bool           { return c.enableDebug }
func (c *Config) RequestTimeout() time.Duration { return c.reqTimeout }
func (c *Config) JwtSecret() string             { return c.jwtSecret }

// ハードコート
func (c *Config) CompressionLevel() int { return compressionLevel }
func (c *Config) JwtExpiredHour() int32 { return jwtExpiredHour }

func (c *Config) Build() *Server {
	e := echo.New()

	// register middlewares
	e.Use(echo_mw.Logger())
	e.Use(echo_mw.Recover())
	e.Use(middleware.Timeout(c.RequestTimeout()))
	e.Use(middleware.Cors(c.Domain()))
	e.Use(middleware.Gzip(c.CompressionLevel()))

	// register route
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

	return &Server{
		c: c,
		e: e,
	}
}

type Server struct {
	c *Config
	e *echo.Echo
}

// see: https://echo.labstack.com/guide/http_server/#http-server
func (r *Server) StartHTTPServer() error {
	return r.e.Start(fmt.Sprintf(":%s", r.c.Port()))
}
