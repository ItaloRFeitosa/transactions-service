package config

import "github.com/italorfeitosa/transactions-service/internal/api/handler"

type Handlers struct {
	Health      *handler.Health
	Transaction *handler.Transaction
	Account     *handler.Account
}

func provideHandlers(c *Container) {
	c.Handlers = new(Handlers)
	c.Handlers.Health = new(handler.Health)
	c.Handlers.Transaction = handler.NewTransaction(c.UseCases.Transaction)
	c.Handlers.Account = handler.NewAccount(c.UseCases.Account)
}
