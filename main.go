package main

import (
	"fmt"
	"github.com/anatofuz/BookmarkTarou/app/config"
	"github.com/anatofuz/BookmarkTarou/web"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	os.Exit(run())
}

func run() int {
	app, err := config.CreateAppComponent()
	if err != nil {
		log.Fatal(err)
		return 1
	}
	fmt.Println(app.UserStore())
	e := echo.New()
	webHandler := web.CreateHandlerImpl()
	webHandler.Perform(e, app.UserStore())
	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
		return 1
	}
	return 0
}
