package usecase

import (
	"errors"
)

var (
	ErrDuplicateEvent = errors.New("duplicate event for account timeslot")
)
