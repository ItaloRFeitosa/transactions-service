package config

import "github.com/italorfeitosa/transactions-service/internal/app"

type UseCases struct {
	Account     app.AccountUseCase
	Transaction app.TransactionUseCase
}

func provideUseCases(c *Container) {
	c.UseCases = new(UseCases)

	c.UseCases.Account = app.NewAccountUseCase(c.Daos.AccountDAO)
	c.UseCases.Transaction = app.NewTransactionUseCase(c.Daos.TransactionDAO, c.Daos.AccountDAO)
}
