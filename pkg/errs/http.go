package errs

import "net/http"

var statusCodeMap = map[Type]int{
	InternalType:     http.StatusInternalServerError,
	NotFoundType:     http.StatusNotFound,
	ValidationType:   http.StatusBadRequest,
	BusinessRuleType: http.StatusUnprocessableEntity,
	ConflictType:     http.StatusConflict,
}

func ToHttpError(err error) HttpError {
	response := HttpError{
		StatusCode: http.StatusInternalServerError,
		Error:      Error{Message: "something went wrong"},
	}
	asError, ok := AsError(err)
	if !ok {
		return response
	}

	code, ok := statusCodeMap[asError.Type]

	if !ok {
		return response
	}

	response.StatusCode = code
	response.Error = asError
	return response
}

type HttpError struct {
	StatusCode int   `json:"-"`
	Error      Error `json:"error"`
}
