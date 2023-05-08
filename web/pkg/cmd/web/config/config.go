package config

import (
	"fmt"
	"os"
	"time"
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
