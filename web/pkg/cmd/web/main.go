package main

import (
	"github.com/labstack/gommon/log"

	"github.com/tatsuyayamauchi/golang-echo-server-example/web/pkg/cmd/web/config"
)

func main() {
	c := config.NewConfig()
	server := c.Build()

	if err := server.StartHTTPServer(); err != nil {
		log.Error(err)
	}
}
