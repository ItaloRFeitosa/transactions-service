package model

// swagger:response
type OpenAccountValidationError struct {
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
			// Example: malformed_open_account_request
			Code string `json:"code"`
			// the error message
			//
			// Example: malformed open account request, check details
			Message string `json:"message"`
			// Details of error
			//
			// Example: ["missing document_type field", "document_number must be a string"]
			Details []string `json:"details"`
		} `json:"error"`
	}
}

// swagger:response
type OpenAccountBusinessRuleError struct {
	// Business Rule Error for open account endpoint
	// in: body
	Body struct {
		// Example: {"type": "business_rule","code": "document_not_allowed", "message": "document type 'SSN' not allowed"}
		// Example: {"type": "business_rule","code": "invalid_document_number", "message": "invalid document number for type 'CPF'"}
		Error struct {
			// business_rule
			Type string `json:"type"`
			// document_not_allowed | invalid_document_number
			Code string `json:"code"`
			// document type 'SSN' not allowed | invalid document number for type 'CPF'
			Message string `json:"message"`
		} `json:"error"`
	}
}

// swagger:response
type AccountNotFoundResponse struct {
	// Business Rule Error for open account endpoint
	// in: body
	Body struct {
		Error struct {
			// not_found
			//
			// Required: true
			// Example: not_found
			Type string `json:"type"`
			// account_not_found
			//
			// Required: true
			// Example: account_not_found
			Code string `json:"code"`
			// the error message
			//
			// Required: true
			// Example: could not found account with id '2154'
			Message string `json:"message"`
		} `json:"error"`
	}
}
