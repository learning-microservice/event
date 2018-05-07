package event

type domainError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e *domainError) Error() string {
	return e.Message
}

type (
	// ErrAlreadyExists is ...
	ErrAlreadyExists struct{ error }
	// ErrAlreadyBooked is ...
	ErrAlreadyBooked struct{ error }
	// ErrBookingFailure is ...
	ErrBookingFailure struct{ error }
)

////////////////////////////////////////////
// Private Functions
////////////////////////////////////////////

func errorAlreadyExists(field string, value interface{}) error {
	return &ErrAlreadyExists{
		&domainError{
			Field:   field,
			Value:   value,
			Message: "event already exists",
		},
	}
}

func errorAlreadyBooked(field string, value interface{}) error {
	return &ErrAlreadyBooked{
		&domainError{
			Field:   field,
			Value:   value,
			Message: "event already booked",
		},
	}
}

func errorBookingFailure(msg string) error {
	return &ErrBookingFailure{
		&domainError{
			Message: msg,
		},
	}
}
