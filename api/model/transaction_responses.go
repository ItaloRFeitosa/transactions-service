package model

// structs only with documentation purpose

// swagger:response
type SaveTransactionValidationError struct {
	// Validation Error for open account endpoint
	// in: body
	Body struct {
		Error struct {
			// the error type
			//
			// Example: validation
			Type string `json:"type"`
			// The error code
			//
			// Example: malformed_save_transaction_request
			Code string `json:"code"`
			// the error message
			//
			// Example: malformed save transaction request, check details
			Message string `json:"message"`
			// Details of error
			//
			// Example: ["missing account_id field", "missing operation_type_id field", "missing amount field"]
			Details []string `json:"details"`
		} `json:"error"`
	}
}

// swagger:response
type SaveTransactionBusinessRuleError struct {
	// Business Rule Error for save transaction endpoint
	// in: body
	Body struct {
		// Example: {"type": "business_rule","code": "account_not_exists", "message": "account with id '6544' not exists"}
		// Example: {"type": "business_rule","code": "invalid_operation_type", "message": "invalid operation type id '8'"}
		// Example: {"type": "business_rule","code": "positive_amount", "message": "amount must be negative for operation type id '1'"}
		// Example: {"type": "business_rule","code": "negative_amount", "message": "amount must be positive for operation type id '4'"}
		Error struct {
			// business_rule
			Type string `json:"type"`
			// account_not_exists | invalid_operation_type | positive_amount | negative_amount
			Code string `json:"code"`
			// account with id '6544' not exists | invalid operation type id '8' | amount must be negative for operation type id '1' | amount must be positive for operation type id '4'
			Message string `json:"message"`
		} `json:"error"`
	}
}
