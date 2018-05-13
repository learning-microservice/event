package events

import (
	"errors"
)

var (
	// common errors
	ErrNotFound        = errors.New("event not found")
	ErrAlreadyModified = errors.New("event has been modified by another user")

	// booking errors
	ErrAlreadyBooked  = errors.New("event has been already booked")
	ErrMaxAssignees   = errors.New("maximum number of assignees is reached")
	ErrBookngDeadline = errors.New("booking deadline has passed")

	// assign errors
	ErrAlreadyAssigned = errors.New("event has been already assigned")
	ErrMaxAttendees    = errors.New("maximum number of attendees is reached")
)
