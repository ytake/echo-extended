package middleware

import (
	echoMiddleware "github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
)

func GlobalMiddleware(e *echo.Echo) {
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
}
