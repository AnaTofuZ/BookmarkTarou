package web

import (
	"github.com/anatofuz/BookmarkTarou/infra/store"
	"github.com/labstack/echo/v4"
)

type WebHandler interface {
	Perform(e *echo.Echo, usrSote store.UserStore)
}

type WebHandlerImpl struct {
}

func CreateWebHandlerImpl() WebHandler {
	return &WebHandlerImpl{}
}

// Perform ルーティングしつつビジネスロジックを呼び出す
func (w *WebHandlerImpl) Perform(e *echo.Echo, usrStore store.UserStore) {
	newUserHandler(e, usrStore)
}
