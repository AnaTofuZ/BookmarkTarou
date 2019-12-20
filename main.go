package main

import (
	"log"
	"os"

	"github.com/anatofuz/BookmarkTarou/app"
)

func main() {
	os.Exit(run())
}

func run() int {
	err := app.Run()
	if err != nil {
		log.Fatal(err)
		return 1
	}
	return 0
}
