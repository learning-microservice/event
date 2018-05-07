package record

import (
	"time"
)

type Booking struct {
	ID         uint64    `gorm:"primary_key"`
	EventID    string    `gorm:"not null"`
	AttendeeID string    `gorm:"not null"`
	BookedAt   time.Time `gorm:"type:datetime;not null"`
}

func (*Booking) TableName() string {
	return "event_bookings"
}
