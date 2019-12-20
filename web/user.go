package web

import (
	"net/http"

	"github.com/anatofuz/BookmarkTarou/infra/store"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	UserStore store.UserStore
}

func newUserHandler(e *echo.Echo, usrStore store.UserStore) {
	uh := &userHandler{
		UserStore: usrStore,
	}
	e.GET("/", uh.hello)
	e.GET("/signup", uh.hello)
}

func (u *userHandler) hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func (u *userHandler) signup(c echo.Context) error {
	return nil
}
