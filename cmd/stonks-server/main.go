package main

import (
	"context"
	"github.com/veremchukvv/stonks-test/internal/api-server"
	"github.com/veremchukvv/stonks-test/internal/config"
	"github.com/veremchukvv/stonks-test/internal/handlers"
	"github.com/veremchukvv/stonks-test/pkg/logging"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	log := logging.NewLogger(false, "console")
	ctx := logging.WithLogger(context.Background(), log)

	log.Info("Starting the app")
	log.Info("Logger initialized ... (1/3)")

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("error occurred when loading configuration: %v", err)
	}
	log.Info("Configuration successfully loaded ... (2/3)")

	handler := new(handlers.Handler)
	srv := api_server.NewServer(cfg.Server.Port, handler.InitRoutes())

	go srv.Run(ctx)

	log.Infof("Server starting on port: %s ... (3/3)", cfg.Server.Port)
	log.Info("App started!")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-interrupt

	log.Info("Shutdown signal received")

	timeout, cancelFunc := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer cancelFunc()

	err = srv.Shutdown(timeout)
	if err != nil {
		log.Errorf("Error when shutdown server: %v", err)
	}
	log.Info("App successfully stopped")
}
