package process

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func GracefulShutdown(shutdownCallback func(context.Context), seconds time.Duration) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit

	slog.Info("starting graceful shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), seconds*time.Second)
	defer cancel()
	defer signal.Stop(quit)

	go shutdownCallback(ctx)

	<-ctx.Done()

	slog.Info("exiting process")
}
