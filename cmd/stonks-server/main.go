package main

import (
	"github.com/veremchukvv/stonks-test/internal/handler"
	"github.com/veremchukvv/stonks-test/internal/api-server"
	"log"
)

func main() {
	srv := new(api_server.Server)
	handlers := new(handler.Handler)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while server is started: %s", err.Error())
	}
}