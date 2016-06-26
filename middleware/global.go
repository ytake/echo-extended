package middleware

import (
	echoMiddleware "github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"echo-extended/foundation"
)

func GlobalMiddleware(e *echo.Echo) {
	e.Use(ApplicationLogger())
	e.Use(echoMiddleware.Recover())
}

func ApplicationLogger() echo.MiddlewareFunc {
	handler := foundation.ApplicationLogFile()
	file := handler.FileHandler
	closure := echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Output: file,
	})
	return closure
}
