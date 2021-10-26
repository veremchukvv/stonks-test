package api_server

import (
	"context"
	"github.com/veremchukvv/stonks-test/pkg/logging"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

//TODO add structure for logger package to inject zap

func NewServer(port string, handler http.Handler) *Server {
	httpServer := &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return &Server{httpServer: httpServer}
}

func (s *Server) Run(ctx context.Context, stopTimeout time.Duration) error {
	logger := logging.FromContext(ctx)
	go func() {
		err := s.httpServer.ListenAndServe()
		if err != nil {
			logger.Fatalf("error starting server %v", err)
		}
	}()
	<-ctx.Done()
	log.Print(ctx)
	ctxstop, cancelstop := context.WithTimeout(context.Background(), stopTimeout)
	defer cancelstop()
	logger.Info("Server is shutting down...")
	if err := s.Shutdown(ctxstop); err != nil {
		logger.Errorf("srv.Shutdown error, %s", err)
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
