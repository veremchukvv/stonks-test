package api_server

import (
	"context"
	"github.com/veremchukvv/stonks-test/pkg/logging"
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

func (s *Server) Run(ctx context.Context) {
	logger := logging.FromContext(ctx)
	//TODO think about making channel for errors
	go func() {
		err := s.httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatalf("error starting server %v", err)
		}
	}()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
