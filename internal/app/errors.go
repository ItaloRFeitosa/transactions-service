package app

import "github.com/italorfeitosa/transactions-service/pkg/errs"

var (
	ErrDocumentNotAllowed = errs.Builder().BusinessRule().
				WithCode("document_not_allowed").
				WithTemplate("document type '%s' not allowed")

	ErrInvalidDocumentNumber = errs.Builder().BusinessRule().
					WithCode("invalid_document_number").
					WithTemplate("invalid document number for type '%s'")

	ErrAccountNotFound = errs.Builder().NotFound().
				WithCode("account_not_found").
				WithTemplate("could not found account with id '%d'")

	ErrAccountNotExists = errs.Builder().BusinessRule().
				WithCode("account_not_exists").
				WithTemplate("account with id '%d' not exists")

	ErrInvalidOperationType = errs.Builder().BusinessRule().
				WithCode("invalid_operation_type").
				WithTemplate("invalid operation type id '%d'")

	ErrPositiveAmount = errs.Builder().BusinessRule().
				WithCode("positive_amount").
				WithTemplate("amount must be negative for operation type id '%d'")

	ErrNegativeAmount = errs.Builder().BusinessRule().
				WithCode("negative_amount").
				WithTemplate("amount must be positive for operation type id '%d'")

	ErrNonPositiveAvailableCreditLimit = errs.Builder().BusinessRule().
						WithCode("non_positive_available_credit_limit").
						WithMessage("available credit limit must be positive")

	ErrNoCreditLimit = errs.Builder().BusinessRule().
				WithCode("no_credit_limit").
				WithTemplate("no credit limit, current limit: %d")

	ErrOldAccountState = errs.Builder().Conflict().
				WithCode("old_account_state").
				WithMessage("old account state, please try again")
)
