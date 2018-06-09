package errors

import (
	"encoding/json"
)

// ApplicationError is ...
type ApplicationError interface {
	error
	Cause() error
	Type() Type
}

////////////////////////////////////////////
// applicationError
////////////////////////////////////////////

type applicationError struct {
	errType Type
	field   string
	value   interface{}
	message string
	cause   error
}

func (e *applicationError) Error() string {
	if len(e.message) > 0 {
		return e.message
	}
	if e.cause != nil {
		return e.cause.Error()
	}
	return "unknown error"
}

func (e *applicationError) Cause() error {
	return e.cause
}

func (e *applicationError) Type() Type {
	return e.errType
}

func (e *applicationError) MarshalJSON() ([]byte, error) {
	var value = e.value
	if v, ok := e.value.(string); ok && len(v) == 0 {
		value = nil
	}
	return json.Marshal(&struct {
		Field   string      `json:"field,omitempty"`
		Value   interface{} `json:"value,omitempty"`
		Message string      `json:"message"`
	}{
		Field:   e.field,
		Value:   value,
		Message: e.Error(),
	})
}

func newApplicationError(errType Type, field string, value interface{}, msg string, cause error) ApplicationError {
	return &applicationError{
		errType: errType,
		field:   field,
		value:   value,
		message: msg,
		cause:   cause,
	}
}

////////////////////////////////////////////
// AlreadyModifiedError
////////////////////////////////////////////

// NewAlreadyModifiedError is ...
func NewAlreadyModifiedError(field string, value interface{}, msg string) error {
	return newApplicationError(AlreadyModifiedErrorType, field, value, msg, nil)
}

////////////////////////////////////////////
// NotFoundError
////////////////////////////////////////////

// NewNotFoundError is ...
func NewNotFoundError(field string, value interface{}, msg string) error {
	return newApplicationError(NotFoundErrorType, field, value, msg, nil)
}

////////////////////////////////////////////
// ValidationError
////////////////////////////////////////////

// NewValidationError is ...
func NewValidationError(field string, value interface{}, msg string) error {
	return newApplicationError(ValidationErrorType, field, value, msg, nil)
}
