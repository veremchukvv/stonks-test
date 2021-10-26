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

//type Service struct {
//	logger *zap.SugaredLogger
//	server *http.Server
//}

func main() {

	//srv := &Service{
	//
	//}

	log := logging.NewLogger(false, "console")
	ctx := logging.WithLogger(context.Background(), log)

	log.Info("Starting the app")
	log.Info("Logger initialized ... (1/3)")

	cfg, err := config.GetConfig()
	log.Infof("timeout is set %v", cfg.Server.ShutdownTimeout)
	if err != nil {
		log.Fatalf("error occured when loading configuration: %v", err)
	}
	log.Info("Configuration successfully loaded ... (2/3)")

	handler := new(handlers.Handler)
	srv := api_server.NewServer(cfg.Server.Port, handler.InitRoutes())


	//ctx, cancel := context.WithCancel(context.Background())

	go srv.Run(ctx, cfg.Server.ShutdownTimeout)

	log.Infof("Server starting on port: %s ... (3/3)", cfg.Server.Port)
	log.Info("App started!")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<- interrupt

	log.Info("Get signal to stop app")

	timeout, cancelFunc := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer cancelFunc()

	err = srv.Shutdown(timeout)
	if err != nil {
		log.Errorf("Error when shutdown server: %v", err)
	}

	log.Info("Server stopped")

	//go func() {
	//	err := srv.Run(cfg.Server.Port, handler.InitRoutes())
	//	if err != nil {
	//		log.Fatalf("error occured while server is started: %s", err.Error())
	//	}
	//}()
	//log.Infof("Server running on port: %s", cfg.Server.Port)

	//go srv.Run(cfg.Server.Port, handler.InitRoutes())
	//go http.ListenAndServe()

	//err = srv.Run(cfg.Server.Port, handler.InitRoutes())
	//log.Infof("Server running on port: %s", cfg.Server.Port)
	//if err := srv.Run(cfg.Server.Port, handler.InitRoutes()); err != nil {
	//	log.Fatalf("error occured while server is started: %s", err.Error())
	//}
}
