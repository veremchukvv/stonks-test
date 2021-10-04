package main

import (
	"github.com/veremchukvv/stonks-test/restapi"
	"log"
)

func main() {
	srv := new(restapi.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while server is started: %s", err.Error())
	}
}