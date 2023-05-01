package handler

import (
	"net/http"
	"net/http/pprof"

	"github.com/labstack/echo/v4"
)

func DebugHandlerFuncs() map[string]echo.HandlerFunc {
	return map[string]echo.HandlerFunc{
		"/cmdline": echo.WrapHandler(http.HandlerFunc(pprof.Cmdline)),
		"/profile": echo.WrapHandler(http.HandlerFunc(pprof.Profile)),
		"/symbol":  echo.WrapHandler(http.HandlerFunc(pprof.Symbol)),
		"/trace":   echo.WrapHandler(http.HandlerFunc(pprof.Trace)),
		"/*":       echo.WrapHandler(http.HandlerFunc(pprof.Index)),
	}
}
