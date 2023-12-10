package model

// structs only with documentation purpose

import "time"

// swagger:model
type OpenAccountRequest struct {
	// Document Type, it can be CPF, CNPJ or NINO
	//
	// Required: true
	// example: CPF
	DocumentType string `json:"document_type"`
	// Document Number, should be valid for given document_type
	//
	// required: true
	// example: 62558363042
	DocumentNumber string `json:"document_number"`
}

// swagger:model
type AccountResponse struct {
	Data Account `json:"data"`
}

// swagger:model
type Account struct {
	// Account ID, identifies account
	// example: 6543
	AccountID int `json:"account_id"`
	// Document Type, it can be CPF, CNPJ or NINO
	// example: CPF
	DocumentType string `json:"document_type"`
	// Document Number, should be valid for given document_type
	// example: 62558363042
	DocumentNumber string    `json:"document_number"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
