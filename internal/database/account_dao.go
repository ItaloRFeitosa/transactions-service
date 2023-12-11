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
	insertAccountQuery = `insert into accounts (document_type, document_number, created_at, updated_at, deleted_at) 
	values (:document_type, :document_number, :created_at, :updated_at, :deleted_at) returning account_id;`

	getAccountQuery    = "select * from accounts where account_id = $1 order by account_id asc limit 1;"
	existsAccountQuery = "select exists (select 1 from accounts where account_id = $1 order by account_id asc limit 1);"
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
		DocumentType:   input.DocumentType,
		DocumentNumber: input.DocumentNumber,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
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

func (dao *accountDAO) Exists(ctx context.Context, accountID int) (bool, error) {
	var exists bool
	err := dao.db.GetContext(ctx, &exists, existsAccountQuery, accountID)

	if err != nil {
		return false, err
	}

	return exists, nil
}
