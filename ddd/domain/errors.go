package domain

type domainError struct {
	Field   string      `json:"field,omitempty"`
	Value   interface{} `json:"value,omitempty"`
	Message string      `json:"message"`
	Cause   error       `json:"-"`
}

func (e *domainError) Error() string {
	if e.Cause != nil {
		return e.Cause.Error()
	}
	if len(e.Message) > 0 {
		return e.Message
	}
	return "unknown"
}

func NotFoundError(field string, value interface{}, msg string, cause error) error {
	return &domainError{
		Field:   field,
		Value:   value,
		Message: msg,
		Cause:   cause,
	}
}

func AlreadyModifiedError(field string, value interface{}, msg string, cause error) error {
	return &domainError{
		Field:   field,
		Value:   value,
		Message: "Record was already modified by another user",
		Cause:   cause,
	}
}
