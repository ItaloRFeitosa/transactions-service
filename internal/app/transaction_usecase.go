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

	account, err := a.accountDAO.Get(ctx, input.AccountID)
	if err != nil {
		return TransactionDTO{}, err
	}

	newCreditLimit := account.AvailableCreditLimit + input.Amount

	if newCreditLimit < 0 {
		return TransactionDTO{}, ErrNoCreditLimit.WithArgs(account.AvailableCreditLimit)
	}

	account.AvailableCreditLimit = newCreditLimit

	return a.accountDAO.SensibilizeTransactionToAccount(ctx, account, input)
}
