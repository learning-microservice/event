package model

import (
	"github.com/learning-microservice/event/ddd/domain/model/shared/account"
	"github.com/learning-microservice/event/ddd/domain/model/shared/event"
)

type Booking struct {
	id         uint
	eventID    event.ID
	attendeeID account.ID
}

func (b *Booking) AttendeeID() account.ID {
	return b.attendeeID
}

func newBooking(eventID event.ID, attendeeID account.ID) Booking {
	return Booking{
		eventID:    eventID,
		attendeeID: attendeeID,
	}
}
