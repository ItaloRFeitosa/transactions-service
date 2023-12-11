package app

import (
	"context"
	"time"
)

type AccountUseCase interface {
	OpenAccount(context.Context, OpenAccountInput) (AccountDTO, error)
	GetAccount(ctx context.Context, accountID int) (AccountDTO, error)
}

type AccountDAO interface {
	Insert(context.Context, OpenAccountInput) (AccountDTO, error)
	Get(ctx context.Context, accountID int) (AccountDTO, error)
	Exists(ctx context.Context, accountID int) (bool, error)
}

type OpenAccountInput struct {
	DocumentType   string
	DocumentNumber string
}

type AccountDTO struct {
	AccountID      int
	DocumentType   string
	DocumentNumber string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
