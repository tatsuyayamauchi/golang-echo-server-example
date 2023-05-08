package main

import (
	"github.com/labstack/gommon/log"

	"github.com/tatsuyayamauchi/golang-echo-server-example/web/pkg/cmd/web/config"
	"github.com/tatsuyayamauchi/golang-echo-server-example/web/pkg/handler"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Error(err)
		panic(err)
	}
	router := handler.NewRouter(c)

	if err := router.StartHTTPServer(); err != nil {
		log.Error(err)
	}
}
