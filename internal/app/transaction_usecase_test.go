package app_test

import (
	"context"
	"testing"
	"time"

	"github.com/italorfeitosa/transactions-service/internal/app"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSaveTransaction(t *testing.T) {
	t.Run("given valid transaction params should return transaction data", func(t *testing.T) {
		// ARRANGE
		var err error

		daos := newFakeDaos()

		accountUseCase := app.NewAccountUseCase(daos.accountDaoFake)
		transactionUseCase := app.NewTransactionUseCase(daos.transactionDaoFake, daos.accountDaoFake)

		openAccountInput := app.OpenAccountInput{
			DocumentType:         app.CPF,
			DocumentNumber:       "05967545310",
			AvailableCreditLimit: 5000,
		}

		newAccount, err := accountUseCase.OpenAccount(context.Background(), openAccountInput)
		require.NoError(t, err)

		saveTransactionInput := app.SaveTransactionInput{
			AccountID:       newAccount.AccountID,
			OperationTypeID: app.PaymentOperation,
			Amount:          2000,
		}

		expectedAvailableCreditLimit := openAccountInput.AvailableCreditLimit + saveTransactionInput.Amount

		// ACT
		transactionDTO, err := transactionUseCase.SaveTransaction(context.Background(), saveTransactionInput)

		// ASSERT
		assert.NoError(t, err)

		currentAccount, err := accountUseCase.GetAccount(context.Background(), newAccount.AccountID)
		assert.NoError(t, err)

		assert.Equal(t, expectedAvailableCreditLimit, currentAccount.AvailableCreditLimit)
		assert.Equal(t, saveTransactionInput.AccountID, transactionDTO.AccountID)
		assert.Equal(t, saveTransactionInput.OperationTypeID, transactionDTO.OperationTypeID)
		assert.Equal(t, saveTransactionInput.Amount, transactionDTO.Amount)
	})

	t.Run("given valid transaction params should return transaction data", func(t *testing.T) {
		// ARRANGE
		var err error

		daos := newFakeDaos()

		accountUseCase := app.NewAccountUseCase(daos.accountDaoFake)
		transactionUseCase := app.NewTransactionUseCase(daos.transactionDaoFake, daos.accountDaoFake)

		openAccountInput := app.OpenAccountInput{
			DocumentType:         app.CPF,
			DocumentNumber:       "05967545310",
			AvailableCreditLimit: 5000,
		}

		newAccount, err := accountUseCase.OpenAccount(context.Background(), openAccountInput)
		require.NoError(t, err)

		saveTransactionInput := app.SaveTransactionInput{
			AccountID:       newAccount.AccountID,
			OperationTypeID: app.PurchaseOperation,
			Amount:          -2000,
		}

		expectedAvailableCreditLimit := openAccountInput.AvailableCreditLimit + saveTransactionInput.Amount

		// ACT
		transactionDTO, err := transactionUseCase.SaveTransaction(context.Background(), saveTransactionInput)

		// ASSERT
		assert.NoError(t, err)

		currentAccount, err := accountUseCase.GetAccount(context.Background(), newAccount.AccountID)
		assert.NoError(t, err)

		assert.Equal(t, expectedAvailableCreditLimit, currentAccount.AvailableCreditLimit)
		assert.Equal(t, saveTransactionInput.AccountID, transactionDTO.AccountID)
		assert.Equal(t, saveTransactionInput.OperationTypeID, transactionDTO.OperationTypeID)
		assert.Equal(t, saveTransactionInput.Amount, transactionDTO.Amount)
	})

	t.Run("given valid transaction params when no credit limit should return error", func(t *testing.T) {
		// ARRANGE
		var err error

		daos := newFakeDaos()

		accountUseCase := app.NewAccountUseCase(daos.accountDaoFake)
		transactionUseCase := app.NewTransactionUseCase(daos.transactionDaoFake, daos.accountDaoFake)

		openAccountInput := app.OpenAccountInput{
			DocumentType:         app.CPF,
			DocumentNumber:       "05967545310",
			AvailableCreditLimit: 5000,
		}

		newAccount, err := accountUseCase.OpenAccount(context.Background(), openAccountInput)
		require.NoError(t, err)

		saveTransactionInput := app.SaveTransactionInput{
			AccountID:       newAccount.AccountID,
			OperationTypeID: app.PurchaseOperation,
			Amount:          -5001,
		}

		// ACT
		_, err = transactionUseCase.SaveTransaction(context.Background(), saveTransactionInput)

		// ASSERT
		assert.EqualError(t, err, app.ErrNoCreditLimit.WithArgs(newAccount.AvailableCreditLimit).Error())
	})

	t.Run("given valid transaction params when account not exists should return error", func(t *testing.T) {
		// ARRANGE
		var err error

		daos := newFakeDaos()

		transactionUseCase := app.NewTransactionUseCase(daos.transactionDaoFake, daos.accountDaoFake)

		saveTransactionInput := app.SaveTransactionInput{
			AccountID:       1,
			OperationTypeID: app.PurchaseOperation,
			Amount:          -5000,
		}

		// ACT
		_, err = transactionUseCase.SaveTransaction(context.Background(), saveTransactionInput)

		// ASSERT
		assert.EqualError(t, err, app.ErrAccountNotExists.WithArgs(saveTransactionInput.AccountID).Error())
	})

	t.Run("given invalid operation type and amount should return error", func(t *testing.T) {
		// ARRANGE
		var err error

		daos := newFakeDaos()

		transactionUseCase := app.NewTransactionUseCase(daos.transactionDaoFake, daos.accountDaoFake)

		saveTransactionInput := app.SaveTransactionInput{
			AccountID:       1,
			OperationTypeID: app.PaymentOperation,
			Amount:          -5000,
		}

		// ACT
		_, err = transactionUseCase.SaveTransaction(context.Background(), saveTransactionInput)

		// ASSERT
		assert.EqualError(t, err, app.ErrNegativeAmount.WithArgs(saveTransactionInput.OperationTypeID).Error())
	})
}

type fakeDaos struct {
	accountDaoFake     *accountDaoFake
	transactionDaoFake *transactionDaoFake
}

func newFakeDaos() *fakeDaos {
	d := new(fakeDaos)
	accounts := make(map[int]app.AccountDTO)
	transactions := make(map[int]app.TransactionDTO)

	d.accountDaoFake = newAccountDaoFake(accounts)
	d.transactionDaoFake = newTransactionDaoFake(accounts, transactions)

	return d
}

type accountDaoFake struct {
	accounts map[int]app.AccountDTO
}

func newAccountDaoFake(accounts map[int]app.AccountDTO) *accountDaoFake {
	return &accountDaoFake{accounts}
}

func (a *accountDaoFake) Insert(ctx context.Context, input app.OpenAccountInput) (app.AccountDTO, error) {
	var accountDTO app.AccountDTO

	accountDTO.AccountID = len(a.accounts) + 1
	accountDTO.AvailableCreditLimit = input.AvailableCreditLimit
	accountDTO.DocumentNumber = input.DocumentNumber
	accountDTO.DocumentType = input.DocumentType
	accountDTO.CreatedAt = time.Now()
	accountDTO.UpdatedAt = time.Now()

	a.accounts[accountDTO.AccountID] = accountDTO

	return accountDTO, nil
}

func (a *accountDaoFake) Get(ctx context.Context, accountID int) (app.AccountDTO, error) {
	foundAcc, ok := a.accounts[accountID]
	if !ok {
		return app.AccountDTO{}, app.ErrAccountNotFound.WithArgs(accountID)
	}
	return foundAcc, nil
}

func (a *accountDaoFake) Exists(ctx context.Context, accountID int) (bool, error) {
	_, ok := a.accounts[accountID]
	return ok, nil
}

type transactionDaoFake struct {
	accounts     map[int]app.AccountDTO
	transactions map[int]app.TransactionDTO
}

func newTransactionDaoFake(accounts map[int]app.AccountDTO, transactions map[int]app.TransactionDTO) *transactionDaoFake {
	return &transactionDaoFake{accounts, transactions}
}

func (t *transactionDaoFake) Insert(ctx context.Context, data app.InsertTransactionData) (app.TransactionDTO, error) {
	var transactionDTO app.TransactionDTO

	foundAccount, ok := t.accounts[data.AccountID]
	if !ok {
		return transactionDTO, app.ErrAccountNotExists.WithArgs(data.AccountID)
	}

	if foundAccount.Version != data.ExpectedAccountVersion {
		return transactionDTO, app.ErrOldAccountState
	}

	foundAccount.AvailableCreditLimit = data.NewAvailableCreditLimit
	foundAccount.Version += 1
	t.accounts[data.AccountID] = foundAccount

	transactionDTO.TransactionID = len(t.transactions) + 1
	transactionDTO.AccountID = data.AccountID
	transactionDTO.OperationTypeID = data.OperationTypeID
	transactionDTO.Amount = data.Amount
	transactionDTO.CreatedAt = time.Now()

	t.transactions[transactionDTO.TransactionID] = transactionDTO

	return transactionDTO, nil
}
