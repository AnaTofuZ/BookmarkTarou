package web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/anatofuz/BookmarkTarou/app"
	"github.com/anatofuz/BookmarkTarou/infra/store"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	App       app.UserApp
	templates map[string]*template.Template
}

func (u *userHandler) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return u.templates[name].ExecuteTemplate(w, "main.tmpl", data)
}

func newUserHandler(e *echo.Echo, usrStore store.UserStore) {
	uh := &userHandler{
		App:       app.NewUserApp(usrStore),
		templates: make(map[string]*template.Template),
	}
	uh.templates["index.tmpl"] = template.Must(template.ParseFiles("templates/index.tmpl", "templates/main.tmpl"))
	uh.templates["signup.tmpl"] = template.Must(template.ParseFiles("templates/signup.tmpl", "templates/main.tmpl"))
	e.Renderer = uh
	e.GET("/", uh.getRoot)
	e.GET("/signup", uh.getSignup)
	e.POST("/signup", uh.postSignup)
}

func (u *userHandler) getRoot(c echo.Context) error {
	return c.Render(http.StatusOK, "index.tmpl", nil)
}

func (u *userHandler) getSignup(c echo.Context) error {
	return c.Render(http.StatusOK, "signup.tmpl", nil)
}

func (u *userHandler) postSignup(c echo.Context) error {
	params := new(struct {
		Name     string `form:"name"`
		Password string `form:"password"`
	})

	err := c.Bind(params)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed signup: %w", err)
	}
	usr, err := u.App.Create(c.Request().Context(), params.Name, params.Password)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed signup: %w", err)
	}
	return c.Render(http.StatusOK, "index.tmpl", map[string]interface{}{
		"User":usr,
	})
}
