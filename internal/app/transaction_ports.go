package app

import (
	"context"
	"time"
)

type TransactionUseCase interface {
	SaveTransaction(context.Context, SaveTransactionInput) (TransactionDTO, error)
}

type TransactionDAO interface {
	Insert(context.Context, SaveTransactionInput) (TransactionDTO, error)
}

type SaveTransactionInput struct {
	AccountID       int
	OperationTypeID int
	Amount          int
}

type TransactionDTO struct {
	TransactionID   int
	AccountID       int
	OperationTypeID int
	Amount          int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
