package main

import (
	"fmt"
	"github.com/anatofuz/BookmarkTarou/app/config"
	"github.com/anatofuz/BookmarkTarou/web"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	os.Exit(run())
}

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("Request Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
}

func run() int {
	app, err := config.CreateAppComponent()
	if err != nil {
		log.Fatal(err)
		return 1
	}
	fmt.Println(app.UserStore())
	e := echo.New()
	e.Use(middleware.BodyDump(bodyDumpHandler))
	webHandler := web.CreateHandlerImpl(web.StoreInterfaces{
		BookmarkStore:    app.BookmarkStore(),
		EntryStore:       app.EntryStore(),
		UserSessionStore: app.UserSessionStore(),
		UserStore:        app.UserStore(),
	})
	webHandler.Perform(e)

	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
		return 1
	}
	return 0
}
