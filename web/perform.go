package web

import (
	"github.com/anatofuz/BookmarkTarou/infra/store"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	Perform(e *echo.Echo, usrSote store.UserStore)
}

type HandlerImpl struct {
}

func CreateHandlerImpl() Handler {
	return &HandlerImpl{}
}

// Perform ルーティングしつつビジネスロジックを呼び出す
func (w *HandlerImpl) Perform(e *echo.Echo, usrStore store.UserStore) {
	newUserHandler(e, usrStore)
}
