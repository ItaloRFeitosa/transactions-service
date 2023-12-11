package database

import (
	"context"
	"time"

	"github.com/italorfeitosa/transactions-service/internal/app"
	"github.com/jmoiron/sqlx"
)

const insertTransactionQuery = `insert into transactions (account_id, operation_type_id, amount, created_at) 
values (:account_id, :operation_type_id, :amount, :created_at) returning transaction_id;`

type transactionDAO struct {
	db *sqlx.DB
}

func NewTransactionDAO(db *sqlx.DB) *transactionDAO {
	return &transactionDAO{db}
}

func (dao *transactionDAO) Insert(ctx context.Context, input app.SaveTransactionInput) (app.TransactionDTO, error) {
	var (
		err            error
		transactionDTO app.TransactionDTO
	)

	transactionModel := TransactionModel{
		AccountID:       input.AccountID,
		OperationTypeID: input.OperationTypeID,
		Amount:          input.Amount,
		CreatedAt:       time.Now(),
	}

	transactionModel.TransactionID, err = insertReturningID(ctx, dao.db, insertTransactionQuery, transactionModel)
	if err != nil {
		return transactionDTO, err
	}

	transactionDTO.TransactionID = transactionModel.TransactionID
	transactionDTO.AccountID = transactionModel.AccountID
	transactionDTO.OperationTypeID = transactionModel.OperationTypeID
	transactionDTO.Amount = transactionModel.Amount
	transactionDTO.CreatedAt = transactionModel.CreatedAt

	return transactionDTO, nil
}
