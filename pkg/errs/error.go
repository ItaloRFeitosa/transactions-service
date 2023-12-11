package errs

import "errors"

type Error struct {
	Type    Type     `json:"type,omitempty"`
	Code    string   `json:"code,omitempty"`
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`

	template string
	err      error
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) Is(target error) bool {
	if targetError, ok := target.(Error); ok {
		return targetError.Code == e.Code && targetError.Type == e.Type
	}

	return errors.Is(e.err, target)
}

func AsError(err error) (Error, bool) {
	asError, ok := err.(Error)

	return asError, ok
}
