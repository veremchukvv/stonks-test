package main

import (
	"github.com/veremchukvv/stonks-test/internal/api-server"
	"github.com/veremchukvv/stonks-test/internal/config"
	"github.com/veremchukvv/stonks-test/internal/handlers"
	"github.com/veremchukvv/stonks-test/pkg/logging"
)

func main() {

	cfg := config.GetConfig()

	logger := logging.NewLogger(false, "console")
	srv := new(api_server.Server)
	handler := new(handlers.Handler)
	logger.Info("Handlers initialized")
	logger.Infof("Server running on port: %s", cfg.Server.Port)
	if err := srv.Run(cfg.Server.Port, handler.InitRoutes()); err != nil {
		logger.Fatalf("error occured while server is started: %s", err.Error())
	}

}