package usecase

import (
	"errors"
)

var (
	ErrNotFound       = errors.New("event not found")
	ErrDuplicateEvent = errors.New("duplicate event for account timeslot")
)
