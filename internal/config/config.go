package config

import "github.com/italorfeitosa/transactions-service/pkg/logger"

// Container is the struct the carries all dependencies of application
type Container struct {
	Env *Env
}

func NewContainer() *Container {
	c := new(Container)

	c.Env = newEnv()

	logger.Init(c.Env.LogLevel)

	return c
}
