package web

import (
	"fmt"
	"net/http"

	"github.com/anatofuz/BookmarkTarou/app"
	"github.com/anatofuz/BookmarkTarou/infra/store"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	App app.UserApp
}

func newUserHandler(e *echo.Echo, usrStore store.UserStore) {
	uh := &userHandler{
		App: app.NewUserApp(usrStore),
	}
	e.GET("/", uh.hello)
	e.GET("/signup", uh.hello)
}

func (u *userHandler) hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func (u *userHandler) signup(c echo.Context) error {
	params := new(struct {
		Name     string `form:name`
		Password string `form:password`
	})

	err := c.Bind(params)
	if err != nil {
		return fmt.Errorf("failed signup: %w", err)
	}
	usr, err := u.App.CreateUser(c.Request().Context(), params.Name, params.Password)
	if err != nil {
		return fmt.Errorf("failed signup: %w", err)
	}
	return c.String(http.StatusOK, usr)
}
