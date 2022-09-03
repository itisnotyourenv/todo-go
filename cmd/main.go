package main

import (
	"github.com/itisnotyourenv/todo-go"
	"github.com/itisnotyourenv/todo-go/pkg/handler"
	"github.com/itisnotyourenv/todo-go/pkg/repository"
	"github.com/itisnotyourenv/todo-go/pkg/service"
	"log"
)

const Port = "8080"

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	log.Println("Starting server on port", Port)
	srv := new(todo.Server)
	if err := srv.Run(Port, handlers.InitRoutes()); err != nil {
		log.Fatalln(err)
	}
}
