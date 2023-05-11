package main

import (
	"github.com/labstack/gommon/log"

	"github.com/tatsuyayamauchi/golang-echo-server-example/web/pkg/cmd/web/config"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Error(err)
		panic(err)
	}
	server := c.Build()

	if err := server.StartHTTPServer(); err != nil {
		log.Error(err)
	}
}
