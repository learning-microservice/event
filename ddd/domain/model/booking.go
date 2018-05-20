package model

import (
	"github.com/learning-microservice/event/ddd/domain/shared/account"
	"github.com/learning-microservice/event/ddd/domain/shared/event"
)

type Booking struct {
	eventID    event.ID
	attendeeID account.ID
}

func (b Booking) AttendeeID() account.ID {
	return b.attendeeID
}

var (
	emptyBooking = Booking{}
)
