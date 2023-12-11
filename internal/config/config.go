package config

import (
	"github.com/italorfeitosa/transactions-service/pkg/logger"
	"github.com/jmoiron/sqlx"
)

// Container is the struct the carries all dependencies of application
type Container struct {
	Env      *Env
	DB       *sqlx.DB
	Daos     *Daos
	UseCases *UseCases
	Handlers *Handlers
}

func NewContainer() *Container {

	c := new(Container)

	provideEnv(c)

	logger.Init(c.Env.LogLevel)

	provideDB(c)
	provideDaos(c)
	provideUseCases(c)
	provideHandlers(c)

	return c
}
