package config

import "github.com/italorfeitosa/transactions-service/pkg/logger"

// Container is the struct the carries all dependencies of application
type Container struct {
	Env      *Env
	Handlers *Handlers
}

func NewContainer() *Container {
	c := new(Container)

	provideEnv(c)

	logger.Init(c.Env.LogLevel)

	provideHandlers(c)

	return c
}
