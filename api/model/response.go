package model

// swagger:model
type HealthResponse struct {
	Data Health `json:"data"`
}

// swagger:model
type Health struct {
	Success bool `json:"success"`
}

// swagger:response
type InternalServerError struct {
	// Internal Server Error
	// in: body
	Body struct {
		Error struct {
			// The validation message
			//
			// Required: true
			// Example: something went wrong, contact support for further details.
			Message string `json:"message"`
		} `json:"error"`
	}
}
