package main

import (
	"github.com/NoahWTeng/todo-api-go/src"
	"log"
)

func main() {
	application, err := app.CreateNewApp()
	if err != nil {
		log.Fatal(err)
	}

	_ = application.Init()
}
