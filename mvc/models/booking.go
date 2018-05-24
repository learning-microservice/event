package models

import (
	"time"

	"github.com/learning-microservice/event/mvc/commons/errors"
	"github.com/learning-microservice/event/mvc/commons/types/account"
	"github.com/learning-microservice/event/mvc/commons/types/event"
)

type Bookings []*Booking

func (bs Bookings) existsAttendee(attendeeID account.ID) bool {
	for _, b := range bs {
		if b.AttendeeID == attendeeID {
			return true
		}
	}
	return false
}

type Booking struct {
	ID         uint       `gorm:"primary_key"`
	EventID    event.ID   `gorm:"not null"`
	AttendeeID account.ID `gorm:"not null"`
	BookedAt   time.Time  `gorm:"type:datetime;not null"`
	OperatorID string     `gorm:"not null"`
}

func (*Booking) TableName() string {
	return "event_bookings"
}

func (e *Event) Book(booking *Booking) error {
	// TODO validate deadline

	// validate already booked
	if e.Bookings.existsAttendee(booking.AttendeeID) {
		return errors.NewValidationError(
			"attendee_id",
			booking.AttendeeID,
			"account already booked",
		)
	}

	// validate max attendees
	if len(e.Bookings) == MAX_ATTENDEES {
		return errors.NewValidationError(
			"attendee_id",
			booking.AttendeeID,
			"maximum number of attendees is reached",
		)
	}

	// set event id
	booking.EventID = e.ID

	// append booking
	e.Bookings = append(e.Bookings, booking)
	return nil
}

const (
	MAX_ATTENDEES = 1
)
