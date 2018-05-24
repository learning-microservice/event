package errors

type Type uint

const (
	ValidationErrorType Type = iota + 1
	NotFoundErrorType
	AlreadyModifiedErrorType
)
