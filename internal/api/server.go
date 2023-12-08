package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/transactions-service/internal/api/response"
	"github.com/italorfeitosa/transactions-service/internal/config"
	"github.com/italorfeitosa/transactions-service/pkg/logger"
	"github.com/italorfeitosa/transactions-service/pkg/middleware"
)

// ListenServer starts server on given port, and returns a function to shutdown server
func ListenServer(container *config.Container) func(context.Context) error {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()

	engine.Use(gin.Recovery())
	engine.Use(middleware.RequestID)
	engine.Use(middleware.Logger)

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Data(gin.H{"ok": true}))
	})

	srv := &http.Server{
		Addr:    container.Env.ServerAddr(),
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("error on listen server", "error", err)
		}
	}()

	return func(ctx context.Context) error {
		return srv.Shutdown(ctx)
	}
}
