package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	"github.com/tatsuyayamauchi/golang-echo-server-example/web/pkg/cmd/web/config"
)

func main() {
	c := config.NewConfig()
	server := c.Build()

	eg, ctx := errgroup.WithContext(context.Background())

	// 複数サーバーが建つことを想定
	eg.Go(func() error {
		return server.StartAPIServer()
	})

	// see https://echo.labstack.com/cookbook/graceful-shutdown/
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-quit:
		// Ctrl+Cなど外部要因で終了した場合はこちら
		fmt.Printf("signal received: %v\n", sig)
		break
	case <-ctx.Done():
		// どれか１つにエラーが発生した場合はこちら
		fmt.Print("ctx.Done() received\n")
		break
	}

	server.Close()

	// サーバー終了以外のエラーが発生していたら通知する
	if err := eg.Wait(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("error occured: %v\n", err)
	}
}
