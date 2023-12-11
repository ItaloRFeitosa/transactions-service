package handler

import "github.com/italorfeitosa/transactions-service/pkg/errs"

var (
	ErrMalformedRequest                = errs.Builder().Validation()
	ErrMalformedOpenAccountRequest     = ErrMalformedRequest.WithCode("malformed_open_account_request")
	ErrMalformedAccountID              = ErrMalformedRequest.WithCode("malformed_account_id").WithMessage("accountID is not an integer")
	ErrMalformedSaveTransactionRequest = ErrMalformedRequest.WithCode("malformed_save_transaction_request")
)
