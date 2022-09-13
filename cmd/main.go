package main

import (
	"github.com/itisnotyourenv/todo-go"
	"github.com/itisnotyourenv/todo-go/pkg/handler"
	"github.com/itisnotyourenv/todo-go/pkg/repository"
	"github.com/itisnotyourenv/todo-go/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	// set format of logs
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// считываем конфиг файл
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	// if you want to use .env file
	// read environments from.env file and write to os.env
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	// connecting to postgres
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASS"),
	})

	if err != nil {
		logrus.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	logrus.Infoln("Starting server on port", viper.GetString("port"))
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalln(err)
	}
}

// initConfig sets the path to the configuration file
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
