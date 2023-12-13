package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/italorfeitosa/transactions-service/internal/app"
	"github.com/jmoiron/sqlx"
)

const (
	insertAccountQuery = `insert into accounts (document_type, document_number, available_credit_limit, created_at, updated_at, deleted_at) 
	values (:document_type, :document_number, :available_credit_limit, :created_at, :updated_at, :deleted_at) returning account_id;`

	getAccountQuery    = "select * from accounts where account_id = $1 order by account_id asc limit 1;"
	existsAccountQuery = "select exists (select 1 from accounts where account_id = $1 order by account_id asc limit 1);"

	updateCreditLimitQuery = `update accounts set available_credit_limit = $1, version = version + 1
	where account_id = $2 and version = $3`
)

type accountDAO struct {
	db *sqlx.DB
}

func NewAccountDAO(db *sqlx.DB) *accountDAO {
	return &accountDAO{db}
}

func (dao *accountDAO) Insert(ctx context.Context, input app.OpenAccountInput) (app.AccountDTO, error) {
	var err error

	accountModel := AccountModel{
		DocumentType:         input.DocumentType,
		DocumentNumber:       input.DocumentNumber,
		AvailableCreditLimit: input.AvailableCreditLimit,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	accountModel.AccountID, err = insertReturningID(ctx, dao.db, insertAccountQuery, accountModel)
	if err != nil {
		return app.AccountDTO{}, err
	}

	return accountModel.ToDTO(), nil
}

func (dao *accountDAO) Get(ctx context.Context, accountID int) (app.AccountDTO, error) {
	var accountModel AccountModel

	err := dao.db.GetContext(ctx, &accountModel, getAccountQuery, accountID)
	if errors.Is(err, sql.ErrNoRows) {
		return app.AccountDTO{}, app.ErrAccountNotFound.WithArgs(accountID)
	}

	if err != nil {
		return app.AccountDTO{}, err
	}

	return accountModel.ToDTO(), nil
}

func (dao *accountDAO) SensibilizeTransactionToAccount(ctx context.Context, account app.AccountDTO, transactionInput app.SaveTransactionInput) (app.TransactionDTO, error) {
	var (
		err            error
		transactionDTO app.TransactionDTO
	)

	tx, err := dao.db.BeginTxx(ctx, &sql.TxOptions{})

	if err != nil {
		return transactionDTO, err
	}

	result, err := tx.ExecContext(ctx, updateCreditLimitQuery, account.AvailableCreditLimit, account.AccountID, account.Version)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return transactionDTO, err
		}
		return transactionDTO, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return transactionDTO, err
	}

	if rowsAffected == 0 {
		return transactionDTO, app.ErrOldAccountState
	}

	transactionModel := TransactionModel{
		AccountID:       transactionInput.AccountID,
		OperationTypeID: transactionInput.OperationTypeID,
		Amount:          transactionInput.Amount,
		CreatedAt:       time.Now(),
	}

	transactionModel.TransactionID, err = insertReturningID(ctx, tx, insertTransactionQuery, transactionModel)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return transactionDTO, err
		}
		return transactionDTO, err
	}

	if err := tx.Commit(); err != nil {
		return transactionDTO, err
	}

	transactionDTO.TransactionID = transactionModel.TransactionID
	transactionDTO.AccountID = transactionModel.AccountID
	transactionDTO.OperationTypeID = transactionModel.OperationTypeID
	transactionDTO.Amount = transactionModel.Amount
	transactionDTO.CreatedAt = transactionModel.CreatedAt

	return transactionDTO, nil
}

func (dao *accountDAO) Exists(ctx context.Context, accountID int) (bool, error) {
	var exists bool
	err := dao.db.GetContext(ctx, &exists, existsAccountQuery, accountID)

	if err != nil {
		return false, err
	}

	return exists, nil
}
