package foundation

import (
	"encoding/json"
	"github.com/labstack/echo"
)

/**
 * for hyper application language(hal+json)
 */
func Hal(context echo.Context, code int, b interface{}) (err error) {
	context.Response().Header().Set(echo.HeaderContentType, ApplicationHalJSON)
	context.Response().WriteHeader(code)
	bytes, _ := json.Marshal(b)
	_, err = context.Response().Write(bytes)
	return
}
