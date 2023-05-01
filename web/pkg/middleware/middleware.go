package middleware

import (
	"fmt"
	"time"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RootMiddleware(domain string, t time.Duration) []echo.MiddlewareFunc {
	// 先頭から適用される
	return []echo.MiddlewareFunc{
		middleware.Logger(),
		middleware.Recover(),
		middleware.TimeoutWithConfig(
			middleware.TimeoutConfig{
				Timeout: t,
			},
		),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{fmt.Sprintf("https://%s", domain)},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}),
		middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
		}),
	}
}

func JwtAuth(secretKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(secretKey),
	})
}
