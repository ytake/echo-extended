package main

import (
	"fmt"
	"github.com/labstack/echo"
	"echo-extended/middleware"
	"echo-extended/foundation"
	"github.com/labstack/echo/engine/standard"
	"github.com/facebookgo/grace/gracehttp"
)

func main() {
	// Echo instance
	e := echo.New()
	applicationConfig := foundation.NewApplicationConfig()
	handler := foundation.ApplicationLogFile()
	foundation.RegisterFileLogger(e, handler)
	// register global middleware
	middleware.GlobalMiddleware(e)
	// register application routes
	foundation.ApplicationRoutes(e)
	defer handler.FileHandler.Close()
	// boot application server
	standard := standard.New(fmt.Sprintf("%s:%s", *applicationConfig.Host, *applicationConfig.Port))
	standard.SetHandler(e)
	gracehttp.Serve(standard.Server)
}
