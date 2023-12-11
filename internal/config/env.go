package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/italorfeitosa/transactions-service/pkg/logger"
)

// Env is the struct that carries all enviroment variables
type Env struct {
	Port     string
	LogLevel string

	DatabaseDSN  string
	DatabaseName string
}

func provideEnv(c *Container) {
	c.Env = new(Env)

	c.Env.Port = "8080"
	if v := os.Getenv("PORT"); v != "" {
		c.Env.Port = v
	}

	c.Env.LogLevel = slog.LevelInfo.String()
	if v := os.Getenv("LOG_LEVEL"); v != "" {
		c.Env.LogLevel = v
	}

	if v := os.Getenv("DATABASE_DSN"); v != "" {
		c.Env.DatabaseDSN = v
	} else {
		logger.Fatal("DATABASE_DSN must be set")
	}

	if v := os.Getenv("DATABASE_NAME"); v != "" {
		c.Env.DatabaseName = v
	} else {
		logger.Fatal("DATABASE_NAME must be set")
	}
}

// ServerAddr returns a string on format of ":<PORT>". E. g ":8080"
func (e *Env) ServerAddr() string {
	return fmt.Sprintf(":%s", e.Port)
}
