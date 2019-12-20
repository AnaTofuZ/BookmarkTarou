package app

import (
	"fmt"

	"github.com/anatofuz/BookmarkTarou/app/config"
	"github.com/anatofuz/BookmarkTarou/web"
	"github.com/labstack/echo/v4"
)

func Run() error {
	app, err := config.CreateAppComponent()
	if err != nil {
		return err
	}
	fmt.Println(app.UserStore())
	e := echo.New()
	webHandler := web.CreateWebHandlerImpl()
	webHandler.Perform(e, app.UserStore())
	return e.Start(":8000")
}
