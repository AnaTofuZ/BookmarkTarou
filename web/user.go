package web

import (
	"fmt"
	"github.com/anatofuz/BookmarkTarou/model"
	"github.com/google/uuid"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/anatofuz/BookmarkTarou/app"
	"github.com/anatofuz/BookmarkTarou/infra/store"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	App       app.UserApp
	templates map[string]*template.Template
	session   store.UserSessionStore
}

func (u *userHandler) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return u.templates[name].ExecuteTemplate(w, "main.tmpl", data)
}

func newUserHandler(e *echo.Echo, usrStore store.UserStore, usrSessionStore store.UserSessionStore) {
	uh := &userHandler{
		App:       app.NewUserApp(usrStore),
		templates: make(map[string]*template.Template),
		session:   usrSessionStore,
	}
	uh.templates["index.tmpl"] = template.Must(template.ParseFiles("templates/index.tmpl", "templates/main.tmpl"))
	uh.templates["signup.tmpl"] = template.Must(template.ParseFiles("templates/signup.tmpl", "templates/main.tmpl"))
	uh.templates["signin.tmpl"] = template.Must(template.ParseFiles("templates/signin.tmpl", "templates/main.tmpl"))
	e.Renderer = uh
	e.GET("/", uh.getRoot)
	e.GET("/signup", uh.getSignup)
	e.GET("/logout", uh.getLogout)
	e.POST("/signup", uh.postSignup)

	e.GET("/signin", uh.getSignin)
	e.POST("/signin", uh.postSignin)
}

type renderVal map[string]interface{}

func (u *userHandler) getRoot(c echo.Context) error {
	usr, _ := u.findUser(c)
	return c.Render(http.StatusOK, "index.tmpl", renderVal{
		"User": usr,
	})
}

func (u *userHandler) getSignup(c echo.Context) error {
	return c.Render(http.StatusOK, "signup.tmpl", nil)
}

func (u *userHandler) getSignin(c echo.Context) error {
	return c.Render(http.StatusOK, "signin.tmpl", nil)
}

func (u *userHandler) postSignup(c echo.Context) error {
	params := new(struct {
		Name     string `form:"name"`
		Password string `form:"password"`
	})

	err := c.Bind(params)
	if err != nil {
		log.Print(err)
		return c.Render(http.StatusOK, "signup.tmpl", renderVal{
			"msg": "ログインダメそう!!!",
		})
	}
	usr, err := u.App.Create(c.Request().Context(), params.Name, params.Password)
	if err != nil {
		log.Print(err)
		return c.Render(http.StatusOK, "signup.tmpl", renderVal{
			"msg": "何らかの影響で作れなかったお!!!",
		})
	}
	uuID := createUUID()
	if err := u.session.Set(uuID, usr); err != nil {
		log.Print(err)
		return c.Render(http.StatusOK, "signup.tmpl", renderVal{
			"msg": "何らかの影響でsessionが作れなかったお!!!",
		})
	}
	c.SetCookie(&http.Cookie{Name: "sessionToken", Value: uuID})
	return c.Redirect(http.StatusSeeOther, "/")
}

func (u *userHandler) postSignin(c echo.Context) error {
	params := new(struct {
		Name     string `form:"name"`
		Password string `form:"password"`
	})

	err := c.Bind(params)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed signin: %w", err)
	}

	usr, err := u.App.SignIn(c.Request().Context(), params.Name, params.Password)
	if err != nil {
		return c.Render(http.StatusOK, "index.tmpl", map[string]interface{}{
			"msg": "ログインに失敗しました...",
		})
	}
	if usr == nil {
		return c.Render(http.StatusOK, "index.tmpl", map[string]interface{}{
			"msg": "ユーザーがいないようです...",
		})
	}

	uuID := createUUID()
	if err := u.session.Set(uuID, usr); err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed save: %w", err)
	}
	c.SetCookie(&http.Cookie{Name: "sessionToken", Value: uuID})

	return c.Redirect(http.StatusSeeOther, "/")
}

func (u *userHandler) findUser(c echo.Context) (*model.User, error) {
	sessionCookie, err := c.Cookie("sessionToken")
	if err != nil {
		return nil, fmt.Errorf("failed faindUser: %w", err)
	}
	if sessionCookie.Value == "" {
		return nil, nil
	}
	uuID := sessionCookie.Value
	usr, err := u.session.Get(uuID)
	if err != nil {
		return nil, fmt.Errorf("failed faindUser: %w", err)
	}
	return usr, nil
}

func (u *userHandler) getLogout(c echo.Context) error {
	sessionCookie, err := c.Cookie("sessionToken")
	if err != nil {
		log.Print("failed faindUser: %w", err)
		return c.Redirect(http.StatusSeeOther, "/")
	}
	if sessionCookie.Value == "" {
		return c.Redirect(http.StatusSeeOther, "/")
	}
	uuID := sessionCookie.Value
	u.session.Remove(uuID)
	return c.Redirect(http.StatusSeeOther, "/")
}

func createUUID() string {
	return uuid.New().String()
}
