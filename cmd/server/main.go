package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/mboufous/berlin-departure-board/internal"
	"github.com/mboufous/berlin-departure-board/internal/api"
	"github.com/mboufous/berlin-departure-board/internal/providers/bvg"
)

func main() {
	// Set up structured logging
	opts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: false,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	logger.Info("Starting Berlin Departure Board service")

	if err := run(); err != nil {
		logger.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}

func run() error {
	// Setup api client
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	registry, err := internal.NewProviderRegistry(
		bvg.NewProvider(httpClient),
	)
	if err != nil {
		return fmt.Errorf("failed to create provider registry: %w", err)
	}

	client, err := internal.NewTransportAPIClient(registry, internal.WithSelectedProvider(bvg.ProviderName))
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}

	// Setup api server
	handler := api.NewHandler(client)
	server := api.NewServer(handler.Routes())
	if err := server.Start(); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
