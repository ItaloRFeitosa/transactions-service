package app

import "context"

type accountUseCase struct {
	accountDAO AccountDAO
}

func NewAccountUseCase(accountDAO AccountDAO) *accountUseCase {
	return &accountUseCase{accountDAO}
}

func (a *accountUseCase) OpenAccount(ctx context.Context, input OpenAccountInput) (AccountDTO, error) {
	if input.AvailableCreditLimit < 1 {
		return AccountDTO{}, ErrNonPositiveAvailableCreditLimit
	}

	if err := ValidateDocument(input.DocumentType, input.DocumentNumber); err != nil {
		return AccountDTO{}, err
	}

	return a.accountDAO.Insert(ctx, input)
}

func (a *accountUseCase) GetAccount(ctx context.Context, accountID int) (AccountDTO, error) {
	return a.accountDAO.Get(ctx, accountID)
}
