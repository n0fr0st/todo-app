package main

import (
	"github.com/n0fr0st/todo-app"
	"github.com/n0fr0st/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocured while running server: %s", err.Error())
	}
}
