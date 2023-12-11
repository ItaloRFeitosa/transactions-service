package database

import (
	"time"

	"github.com/italorfeitosa/transactions-service/internal/app"
)

type AccountModel struct {
	AccountID      int        `db:"account_id"`
	DocumentType   string     `db:"document_type"`
	DocumentNumber string     `db:"document_number"`
	CreatedAt      time.Time  `db:"created_at"`
	UpdatedAt      time.Time  `db:"updated_at"`
	DeletedAt      *time.Time `db:"deleted_at"`
}

func (a AccountModel) ToDTO() app.AccountDTO {
	var accountDTO app.AccountDTO

	accountDTO.AccountID = a.AccountID
	accountDTO.DocumentNumber = a.DocumentNumber
	accountDTO.DocumentType = a.DocumentType
	accountDTO.CreatedAt = a.CreatedAt
	accountDTO.UpdatedAt = a.UpdatedAt

	return accountDTO
}
