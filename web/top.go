package web

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func TopRender(e *echo.Echo) {
	e.GET("/", hello)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}
