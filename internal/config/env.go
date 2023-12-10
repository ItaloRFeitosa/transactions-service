package config

import (
	"fmt"
	"log/slog"
	"os"
)

// Env is the struct that carries all enviroment variables
type Env struct {
	Port     string
	LogLevel string
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
}

// ServerAddr returns a string on format of ":<PORT>". E. g ":8080"
func (e *Env) ServerAddr() string {
	return fmt.Sprintf(":%s", e.Port)
}
