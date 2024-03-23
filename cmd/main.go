package main

import (
	"log"
	"github.com/Jereyji/FQW.git"
	"github.com/Jereyji/FQW.git/pkg/handler"
)

func main() {
	handler := new(handler.Handler)
	srv := new(todo.Server)

	if err := srv.Run("8000", handler.InitRoutes()); err !=  nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
	
}
