package model

import "time"

// swagger:model
type SaveTransactionRequest struct {
	// Account ID
	//
	// Required: true
	// Example: 6541
	AccountID int `json:"account_id"`
	// Operation Type ID: 1 (Purchase) - 2 (Purchase In Installments) - 3 (Withdraw) - 4(Payment)
	//
	// Required: true
	// Example: 1
	OperationTypeID int `json:"operation_type_id"`
	// Amount in cents must be positive when OperationTypeID == 4, otherwise is negative
	//
	// Required: true
	// Example: -20000
	Amount int `json:"amount"`
}

// swagger:model
type TransactionResponse struct {
	Data Transaction `json:"data"`
}

// swagger:model
type Transaction struct {
	TransactionID   int       `json:"transaction_id"`
	AccountID       int       `json:"account_id"`
	OperationTypeID int       `json:"operation_type_id"`
	Amount          int       `json:"amount"`
	CreatedAt       time.Time `json:"created_at"`
}
