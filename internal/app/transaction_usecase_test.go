package app_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/italorfeitosa/transactions-service/internal/app"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type accountDaoFake struct {
	accounts map[int]app.AccountDTO
}

func newAccountDaoFake() *accountDaoFake {
	return &accountDaoFake{accounts: make(map[int]app.AccountDTO)}
}
func (a *accountDaoFake) Insert(ctx context.Context, input app.OpenAccountInput) (app.AccountDTO, error) {
	var accountDTO app.AccountDTO

	accountDTO.AccountID = len(a.accounts) + 1
	accountDTO.AvailableCreditLimit = 5000

	a.accounts[accountDTO.AccountID] = accountDTO

	return accountDTO, nil
}

func (a *accountDaoFake) SensibilizeTransactionToAccount(ctx context.Context, acc app.AccountDTO, transac app.SaveTransactionInput) (app.TransactionDTO, error) {
	foundAcc, ok := a.accounts[acc.AccountID]
	if !ok {
		return app.TransactionDTO{}, fmt.Errorf("account not found")
	}

	foundAcc.AvailableCreditLimit = acc.AvailableCreditLimit
	a.accounts[acc.AccountID] = foundAcc
	return app.TransactionDTO{}, nil
}

func (a *accountDaoFake) Get(ctx context.Context, accountID int) (app.AccountDTO, error) {
	foundAcc, ok := a.accounts[accountID]
	if !ok {
		return app.AccountDTO{}, fmt.Errorf("account not found")
	}
	return foundAcc, nil
}
func (a *accountDaoFake) Exists(ctx context.Context, accountID int) (bool, error) {
	return true, nil
}

func TestSaveTransaction(t *testing.T) {
	t.Run("given valid transaction params should return transaction data", func(t *testing.T) {
		var err error

		accountDao := newAccountDaoFake()

		newAccount, err := accountDao.Insert(context.Background(), app.OpenAccountInput{
			AvailableCreditLimit: 5000,
		})
		require.NoError(t, err)

		usecase := app.NewTransactionUseCase(nil, accountDao)

		_, err = usecase.SaveTransaction(context.Background(), app.SaveTransactionInput{
			AccountID:       newAccount.AccountID,
			OperationTypeID: app.PaymentOperation,
			Amount:          2000,
		})

		assert.NoError(t, err)
		currentAccount, err := accountDao.Get(context.Background(), newAccount.AccountID)
		assert.NoError(t, err)
		assert.Equal(t, 7000, currentAccount.AvailableCreditLimit)
	})

	t.Run("given valid transaction params should return transaction data", func(t *testing.T) {
		var err error

		accountDao := newAccountDaoFake()

		newAccount, err := accountDao.Insert(context.Background(), app.OpenAccountInput{
			AvailableCreditLimit: 5000,
		})
		require.NoError(t, err)

		usecase := app.NewTransactionUseCase(nil, accountDao)

		_, err = usecase.SaveTransaction(context.Background(), app.SaveTransactionInput{
			AccountID:       newAccount.AccountID,
			OperationTypeID: app.PurchaseOperation,
			Amount:          -2000,
		})

		assert.NoError(t, err)
		currentAccount, err := accountDao.Get(context.Background(), newAccount.AccountID)
		assert.NoError(t, err)
		assert.Equal(t, 3000, currentAccount.AvailableCreditLimit)
	})

	t.Run("given valid transaction params when no credit limit should return error", func(t *testing.T) {
		var err error

		accountDao := newAccountDaoFake()

		newAccount, err := accountDao.Insert(context.Background(), app.OpenAccountInput{
			AvailableCreditLimit: 5000,
		})
		require.NoError(t, err)

		usecase := app.NewTransactionUseCase(nil, accountDao)

		_, err = usecase.SaveTransaction(context.Background(), app.SaveTransactionInput{
			AccountID:       newAccount.AccountID,
			OperationTypeID: app.PurchaseOperation,
			Amount:          -5001,
		})

		assert.ErrorContains(t, err, "no credit limit")
	})
}
