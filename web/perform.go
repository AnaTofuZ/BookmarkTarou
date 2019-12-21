package web

import (
	"fmt"
	"github.com/anatofuz/BookmarkTarou/model"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/anatofuz/BookmarkTarou/app"
	"github.com/anatofuz/BookmarkTarou/infra/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler interface {
	Perform(e *echo.Echo)
}

type HandlerImpl struct {
	UserApp     app.UserApp
	BookmarkApp app.BookmarkApp
	templates   map[string]*template.Template
	session     store.UserSessionStore
}

func CreateHandlerImpl(storeInteraces StoreInterfaces) Handler {
	return &HandlerImpl{
		UserApp:     app.NewUserApp(storeInteraces.UserStore),
		BookmarkApp: app.NewBookmarkApp(storeInteraces.EntryStore, storeInteraces.BookmarkStore),
		templates:   make(map[string]*template.Template),
		session:     storeInteraces.UserSessionStore,
	}
}

func (h *HandlerImpl) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return h.templates[name].ExecuteTemplate(w, "main.tmpl", data)
}

func (uh *HandlerImpl) Perform(e *echo.Echo) {
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(middleware.Logger())
	uh.templates["index.tmpl"] = template.Must(template.ParseFiles("templates/index.tmpl", "templates/main.tmpl"))
	uh.templates["signup.tmpl"] = template.Must(template.ParseFiles("templates/signup.tmpl", "templates/main.tmpl"))
	uh.templates["signin.tmpl"] = template.Must(template.ParseFiles("templates/signin.tmpl", "templates/main.tmpl"))
	uh.templates["add_bookmark.tmpl"] = template.Must(template.ParseFiles("templates/add_bookmark.tmpl", "templates/main.tmpl"))
	e.Renderer = uh
	e.GET("/", uh.getRoot)
	e.GET("/signup", uh.getSignup)
	e.GET("/logout", uh.getLogout)
	e.POST("/signup", uh.postSignup)

	e.GET("/signin", uh.getSignin)
	e.POST("/signin", uh.postSignin)

	e.GET("/bookmark/add", uh.getbookmarkAdd)
	e.POST("/bookmark/add", uh.postbookmarkAdd)
}

type renderVal map[string]interface{}

func (u *HandlerImpl) getRoot(c echo.Context) error {
	usr, _ := u.findUser(c)
	lwbc, _ := u.BookmarkApp.ListWithBCount(c.Request().Context())
	return c.Render(http.StatusOK, "index.tmpl", renderVal{
		"User":    usr,
		"Entries": *lwbc,
	})
}

func (u *HandlerImpl) getSignup(c echo.Context) error {
	return c.Render(http.StatusOK, "signup.tmpl", nil)
}

func (u *HandlerImpl) getSignin(c echo.Context) error {
	return c.Render(http.StatusOK, "signin.tmpl", nil)
}

func (u *HandlerImpl) postSignup(c echo.Context) error {
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
	usr, err := u.UserApp.Create(c.Request().Context(), params.Name, params.Password)
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

func (u *HandlerImpl) postSignin(c echo.Context) error {
	params := new(struct {
		Name     string `form:"name"`
		Password string `form:"password"`
	})

	err := c.Bind(params)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed signin: %w", err)
	}

	usr, err := u.UserApp.SignIn(c.Request().Context(), params.Name, params.Password)
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

func (u *HandlerImpl) findUser(c echo.Context) (*model.User, error) {
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

func (u *HandlerImpl) getLogout(c echo.Context) error {
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

func (h *HandlerImpl) getbookmarkAdd(c echo.Context) error {
	usr, err := h.findUser(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/")
		return c.Render(http.StatusOK, "/", renderVal{
			"msg": "ログインしてからで!!",
		})
	}
	if usr == nil {
		c.Redirect(http.StatusSeeOther, "/")
		return c.Render(http.StatusOK, "/", renderVal{
			"msg": "ログインしてからで!!",
		})
	}
	return c.Render(http.StatusOK, "add_bookmark.tmpl", nil)
}

func (h *HandlerImpl) postbookmarkAdd(c echo.Context) error {
	usr, err := h.findUser(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/")
		return c.Render(http.StatusOK, "/", renderVal{
			"msg": "ログインしてからで!!",
		})
	}

	params := new(struct {
		URL     string `form:"url"`
		Comment string `form:"comment"`
	})

	err = c.Bind(params)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed postBookmark: %w", err)
	}

	entry, err := h.BookmarkApp.Create(c.Request().Context(), usr.ID, params.URL, params.Comment)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed postBookmark: %w", err)
	}
	fmt.Println(entry)
	return c.Redirect(http.StatusSeeOther, "/")
}

func createUUID() string {
	return uuid.New().String()
}
