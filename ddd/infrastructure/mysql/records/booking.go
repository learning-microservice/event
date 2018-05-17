package records

import (
	"time"
)

type Booking struct {
	ID         uint      `gorm:"primary_key"`
	EventID    uint      `gorm:"not null"`
	AttendeeID string    `gorm:"not null"`
	BookedAt   time.Time `gorm:"type:datetime;not null"`
	IsDeleted  bool      `gorm:"-"`
}

func (*Booking) TableName() string {
	return "event_bookings"
}
