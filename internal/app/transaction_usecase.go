package app

import "context"

type transactionUseCase struct {
	transactionDAO TransactionDAO
	accountDAO     AccountDAO
}

func NewTransactionUseCase(transactionDAO TransactionDAO, accountDAO AccountDAO) *transactionUseCase {
	return &transactionUseCase{transactionDAO, accountDAO}
}

func (a *transactionUseCase) SaveTransaction(ctx context.Context, input SaveTransactionInput) (TransactionDTO, error) {
	if err := ValidateOperationType(input.OperationTypeID, input.Amount); err != nil {
		return TransactionDTO{}, err
	}

	exists, err := a.accountDAO.Exists(ctx, input.AccountID)
	if err != nil {
		return TransactionDTO{}, err
	}

	if !exists {
		return TransactionDTO{}, ErrAccountNotExists.WithArgs(input.AccountID)
	}

	return a.transactionDAO.Insert(ctx, input)
}
