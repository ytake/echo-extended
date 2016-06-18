package main

import (
	"github.com/labstack/echo"
	"echo-extended/middleware"
	"github.com/labstack/echo/engine/fasthttp"
)

func main() {
	// Echo instance
	e := echo.New()
	// register global middleware
	middleware.GlobalMiddleware(e)
	// boot application server
	e.Run(fasthttp.New(":8080"))
}
