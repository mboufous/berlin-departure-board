package api

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	defaultPort     = "8080"
	shutdownTimeout = 5 * time.Second
	writeTimeout    = 30 * time.Second
	readTimeout     = 15 * time.Second
	idleTimeout     = 60 * time.Second
)

type Server struct {
	server *http.Server
}

func NewServer(handler http.Handler) *Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	return &Server{
		server: &http.Server{
			Addr:         ":" + port,
			Handler:      handler,
			WriteTimeout: writeTimeout,
			ReadTimeout:  readTimeout,
			IdleTimeout:  idleTimeout,
		},
	}
}

func (s *Server) Start() error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	serverErrors := make(chan error, 1)

	go func() {
		slog.Info("starting server", "addr", s.server.Addr)
		serverErrors <- s.server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("server error: %w", err)
		}
	case <-stop:
		slog.Info("received shutdown signal")
		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := s.server.Shutdown(ctx); err != nil {
			return fmt.Errorf("failed to shutdown server: %w", err)
		}
		slog.Info("server gracefully stopped")
	}
	return nil
}
