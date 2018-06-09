package errors

// Type is ...
type Type uint

// Error Types
const (
	ValidationErrorType Type = iota + 1
	NotFoundErrorType
	AlreadyModifiedErrorType
)
