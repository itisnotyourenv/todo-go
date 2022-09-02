package main

import (
	"github.com/itisnotyourenv/todo-go"
	"github.com/itisnotyourenv/todo-go/pkg/handler"
	"log"
)

const Port = "8080"

func main() {
	log.Println("Starting server on port", Port)
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run(Port, handlers.InitRoutes()); err != nil {
		log.Fatalln(err)
	}
}
