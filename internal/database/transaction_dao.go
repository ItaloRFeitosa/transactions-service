package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/italorfeitosa/transactions-service/internal/app"
	"github.com/jmoiron/sqlx"
)

const (
	insertTransactionQuery = `insert into transactions (account_id, operation_type_id, amount, created_at) 
values (:account_id, :operation_type_id, :amount, :created_at) returning transaction_id;`

	updateCreditLimitQuery = `update accounts set available_credit_limit = $1, version = version + 1
where account_id = $2 and version = $3`
)

type transactionDAO struct {
	db *sqlx.DB
}

func NewTransactionDAO(db *sqlx.DB) *transactionDAO {
	return &transactionDAO{db}
}

func (dao *transactionDAO) Insert(ctx context.Context, transactionData app.InsertTransactionData) (app.TransactionDTO, error) {
	var (
		err            error
		transactionDTO app.TransactionDTO
	)

	tx, err := dao.db.BeginTxx(ctx, &sql.TxOptions{})

	if err != nil {
		return transactionDTO, err
	}

	result, err := tx.ExecContext(ctx, updateCreditLimitQuery, transactionData.NewAvailableCreditLimit, transactionData.AccountID, transactionData.ExpectedAccountVersion)
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
		AccountID:       transactionData.AccountID,
		OperationTypeID: transactionData.OperationTypeID,
		Amount:          transactionData.Amount,
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
