package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/veremchukvv/stonks-test/internal/apiserver"
	"github.com/veremchukvv/stonks-test/internal/config"
	"github.com/veremchukvv/stonks-test/internal/handlers"
	"github.com/veremchukvv/stonks-test/internal/repository"
	"github.com/veremchukvv/stonks-test/internal/repository/pg"
	"github.com/veremchukvv/stonks-test/internal/service"
	"github.com/veremchukvv/stonks-test/pkg/hash"
	"github.com/veremchukvv/stonks-test/pkg/logging"
)

func main() {
	// TODO display error on main page when backend is unavailable

	log := logging.NewLogger(false, "console")
	ctx := logging.WithLogger(context.Background(), log)

	log.Info("Starting the app")
	log.Info("Logger initialized ... (1/4)")

	e := os.Getenv("IS_PRODUCTION")
	if e == "" {
		log.Info("app started in dev environment")
	} else {
		log.Info("app started in production environment")
	}

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("error occurred when loading configuration: %v", err)
	}
	log.Info("Configuration successfully loaded ... (2/4)")

	log.Info(cfg.Server.CORS)

	db, err := pg.NewPG(ctx, cfg.DB.URL)
	if err != nil {
		log.Fatalf("Can't connect to database %v", err)
	}

	err = pg.HealthPG(ctx, db)
	if err != nil {
		log.Infof("db healtchek error: %v", err)
	}

	log.Info("Database connection OK ... (3/4)")

	repo := repository.NewStore(ctx, db)
	hasher := hash.NewBCPasswordHasher(ctx)
	services := service.NewService(repo, hasher)
	handler := handlers.NewHandlers(ctx, services)
	srv := apiserver.NewServer(cfg.Server.Port, handler.InitRoutes())

	go srv.Run(ctx)

	log.Infof("Server starting on port: %s ... (4/4)", cfg.Server.Port)

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
