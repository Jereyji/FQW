package main

import (
	"log"

	"github.com/Jereyji/FQW.git"
	"github.com/Jereyji/FQW.git/pkg/handler"
	"github.com/Jereyji/FQW.git/pkg/repository"
	"github.com/Jereyji/FQW.git/pkg/service"
)

func main() {
	repos := repository.NewReposizitory()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err !=  nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
	
}
