package models

import (
	"time"

	"github.com/learning-microservice/event/mvc/commons/errors"
	"github.com/learning-microservice/event/mvc/commons/types/account"
)

// Cancel is ...
type Cancel struct {
	ID         uint      `gorm:"primary_key"`
	BookingID  uint      `gorm:"not null"`
	CanceledAt time.Time `gorm:"type:datetime;not null"`
	Reason     string    `gorm:"not null"`
	OperatorID string    `gorm:"not null"`
}

// TableName is ...
func (*Cancel) TableName() string {
	return "event_cancels"
}

// Cancel is ...
func (e *Event) Cancel(attendeeID account.ID, cancel *Cancel) error {
	booking := e.Bookings.findAttendee(attendeeID)
	if booking == nil {
		return errors.NewValidationError(
			"attendee_id",
			attendeeID,
			"not booked event",
		)
	}

	var filteredBookings Bookings
	{
		for _, b := range e.Bookings {
			if b.AttendeeID != attendeeID {
				filteredBookings = append(filteredBookings, b)
			}
		}
		e.Bookings = filteredBookings
	}

	cancel.BookingID = booking.ID

	return nil
}
