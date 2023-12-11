package main

import (
	"context"
	"log/slog"

	"github.com/italorfeitosa/transactions-service/internal/api"
	"github.com/italorfeitosa/transactions-service/internal/config"
	"github.com/italorfeitosa/transactions-service/internal/database"
	"github.com/italorfeitosa/transactions-service/pkg/process"
)

func main() {
	container := config.NewContainer()

	slog.Info("starting server")
	shutdownServer := api.ListenServer(container)

	process.GracefulShutdown(func(ctx context.Context) {
		slog.Info("server shutdown in progress")
		if err := shutdownServer(ctx); err != nil {
			slog.Error("couldn't shutdown server", "error", err)
		}
		slog.Info("server is shutdown")
		slog.Info("database close in progress")
		if err := database.Close(container.DB); err != nil {
			slog.Error("couldn't close database", "error", err)
		}
		slog.Info("database was closed")
	}, 5)
}
