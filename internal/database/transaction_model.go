package database

import (
	"time"
)

type TransactionModel struct {
	TransactionID   int       `db:"transaction_id"`
	AccountID       int       `db:"account_id"`
	OperationTypeID int       `db:"operation_type_id"`
	Amount          int       `db:"amount"`
	CreatedAt       time.Time `db:"created_at"`
}
