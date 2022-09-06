package main

import (
	"github.com/itisnotyourenv/todo-go"
	"github.com/itisnotyourenv/todo-go/pkg/handler"
	"github.com/itisnotyourenv/todo-go/pkg/repository"
	"github.com/itisnotyourenv/todo-go/pkg/service"
	"github.com/spf13/viper"
	"log"
)

const Port = "8080"

func main() {
	// считываем конфиг файл
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	//createMigrate()
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	log.Println("Starting server on port", viper.GetString("port"))
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalln(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
