package main

import (
	_ "github.com/lib/pq"
	"log"
	"github.com/Jereyji/FQW.git"
	"github.com/Jereyji/FQW.git/internal/handler"
	"github.com/Jereyji/FQW.git/internal/repository"
	"github.com/Jereyji/FQW.git/internal/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing configs: %s", err.Error())
	}

	db, err := repository.newPostgresDB(repository.Config{
		Host: "localhost",
		Port: "5436",
		Username: "postgres",
		Password: "qwerty",
		DBName: "postgres",
		SSLMode: "disable",
	})
	if err != nil {
		log.Fatalf("Fail to initialize database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
