package web

import (
	"github.com/anatofuz/BookmarkTarou/infra/store"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler interface {
	Perform(e *echo.Echo, usrSote store.UserStore, usrSessionStore store.UserSessionStore)
}

type HandlerImpl struct {
}

func CreateHandlerImpl() Handler {
	return &HandlerImpl{}
}

// Perform ルーティングしつつビジネスロジックを呼び出す
func (w *HandlerImpl) Perform(e *echo.Echo, usrStore store.UserStore, usrSessionStore store.UserSessionStore) {
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(middleware.Logger())
	newUserHandler(e, usrStore, usrSessionStore)
}
