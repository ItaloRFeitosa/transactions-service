package config

import (
	"github.com/italorfeitosa/transactions-service/internal/app"
	"github.com/italorfeitosa/transactions-service/internal/database"
	"github.com/italorfeitosa/transactions-service/pkg/logger"
)

func provideDB(c *Container) {
	var err error
	c.DB, err = database.Init(c.Env.DatabaseDSN)
	if err != nil {
		logger.Fatal("could not init DB instance", "error", err)
	}

	err = database.MigrateUp(c.DB, c.Env.DatabaseName)
	if err != nil {
		logger.Fatal("could not run migrations", "error", err)
	}
}

type Daos struct {
	AccountDAO     app.AccountDAO
	TransactionDAO app.TransactionDAO
}

func provideDaos(c *Container) {
	c.Daos = new(Daos)

	c.Daos.AccountDAO = database.NewAccountDAO(c.DB)
	c.Daos.TransactionDAO = database.NewTransactionDAO(c.DB)
}
