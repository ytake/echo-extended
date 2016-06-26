package foundation

import (
	"net/http"
	"github.com/labstack/echo"
)

func ApplicationRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return Hal(c, http.StatusOK, []string{})
	})
	/*
	e.Get("/uri", action.YourAction(), // middleware.RouteMiddleware())
	*/
	// or
	/*
	v1 := e.Group("/v1")
	{
		v1.Get("/uri", action.YourAction(), // middleware.RouteMiddleware())
	}
	*/
}
