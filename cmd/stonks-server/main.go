package main

import (
	"github.com/veremchukvv/stonks-test/internal/api-server"
	"github.com/veremchukvv/stonks-test/internal/handlers"
	"github.com/veremchukvv/stonks-test/pkg/logging"
	"go.uber.org/zap"
)

func main() {
	logger := logging.NewLogger(false, "console")
	zap.ReplaceGlobals(logger)
	srv := new(api_server.Server)
	handler := new(handlers.Handler)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		zap.S().Fatalf("error occured while server is started: %s", err.Error())
	}
}