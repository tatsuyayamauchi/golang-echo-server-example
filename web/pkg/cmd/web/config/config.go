package config

import (
	"time"

	"github.com/labstack/echo/v4"
	echo_mw "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/pflag"

	"github.com/tatsuyayamauchi/golang-echo-server-example/web/pkg/handler"
	"github.com/tatsuyayamauchi/golang-echo-server-example/web/pkg/middleware"
)

const (
	dummyJwtSecretKey       string        = "02166ee6e7e830a70f0df089cffb2eff65db7f87c9156ba5e4a1c116f15f5c2e124577a3e2a49a37dd3c32f8b8970ff76b1e0b379f49994e0150a75d8153552e"
	defaultTimeout          time.Duration = 10 * time.Second
	defaultShutdownDuration time.Duration = 1 * time.Second

	compressionLevel int   = 5
	jwtExpiredHour   int32 = 72
)

type Config struct {
	domain     string
	addr       string
	reqTimeout time.Duration
	jwtSecret  string

	enableDebug bool
}

func NewConfig() *Config {
	c := &Config{}

	pflag.StringVar(&c.domain, "domain", "localhost", "APIサーバで使用されるドメインです")
	pflag.StringVar(&c.addr, "addr", ":8080", "APIサーバで使用されるアドレスです")
	pflag.StringVar(&c.jwtSecret, "jwt-secret-key", dummyJwtSecretKey, "JWTで使用する秘密鍵です")
	pflag.DurationVar(&c.reqTimeout, "timeout-sec", defaultTimeout, "APIサーバのタイムアウト(秒)です")

	pflag.BoolVar(&c.enableDebug, "debug", false, "デバッグモードで起動するかどうかを選択します")

	pflag.Parse()
	return c
}

func (c *Config) Build() *Server {
	e := echo.New()

	// register middlewares
	e.Use(echo_mw.Logger())
	e.Use(echo_mw.Recover())
	e.Use(middleware.Timeout(c.reqTimeout))
	e.Use(middleware.Cors(c.domain))
	e.Use(middleware.Gzip(compressionLevel))

	// register route
	loginGroup := e.Group("/v1/login")
	loginGroup.POST("", handler.LoginHandlerFunc(jwtExpiredHour, c.jwtSecret))

	v1Group := e.Group("/v1")
	v1Group.Use(middleware.JwtAuth(c.jwtSecret))
	v1Group.GET("/hello", handler.HelloHandlerFunc("v1/hello handler"))
	v1Group.GET("/hello2", handler.HelloHandlerFunc("v1/hello2 handler"))

	if c.enableDebug {
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
func (r *Server) StartAPIServer() error {
	return r.e.Start(r.c.addr)
}

func (r *Server) Close() {
	// TODO: ヘルスチェックを失敗させるようにしたい

	if r.e != nil {
		_ = r.e.Close()
	}
}
