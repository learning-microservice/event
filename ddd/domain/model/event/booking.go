package event

import (
	"time"
)

////////////////////////////////////////////
// Bookings
////////////////////////////////////////////

type Bookings []Booking

func (bs *Bookings) add(id ID, attendeeID AttendeeID) error {
	for _, book := range *bs {
		if book.attendeeID == attendeeID {
			return errorBookingFailure("event already booked")
		}
	}
	*bs = append(*bs, Booking{
		eventID:    id,
		attendeeID: attendeeID,
	})
	return nil
}

////////////////////////////////////////////
// Booking
////////////////////////////////////////////

type Booking struct {
	id         uint64
	eventID    ID
	attendeeID AttendeeID
	bookedAt   time.Time
	isCanceled bool
}

func (b *Booking) ID() uint64 {
	return b.id
}

func (b *Booking) EventID() ID {
	return b.eventID
}

func (b *Booking) AttendeeID() AttendeeID {
	return b.attendeeID
}

func (b *Booking) IsNew() bool {
	return b.id == 0 && !b.isCanceled
}

func (b *Booking) IsCanceled() bool {
	return b.id > 0 && b.isCanceled
}
