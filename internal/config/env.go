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

func newEnv() *Env {
	e := new(Env)

	e.Port = "8080"
	if v := os.Getenv("PORT"); v != "" {
		e.Port = v
	}

	e.LogLevel = slog.LevelInfo.String()
	if v := os.Getenv("LOG_LEVEL"); v != "" {
		e.LogLevel = v
	}

	return e
}

// ServerAddr returns a string on format of ":<PORT>". E. g ":8080"
func (e *Env) ServerAddr() string {
	return fmt.Sprintf(":%s", e.Port)
}
