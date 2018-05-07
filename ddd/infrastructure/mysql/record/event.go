package record

import (
	"time"
)

type Event struct {
	ID                string    `gorm:"primary_key"`
	Category          string    `gorm:"not null"`
	Contents          []byte    `gorm:"type:json"`
	StartAt           time.Time `gorm:"type:datetime;not null"`
	EndAt             time.Time `gorm:"type:datetime;not null"`
	IsPrivated        bool      `gorm:"not null"`
	PublishedAt       time.Time `gorm:"type:datetime;not null"`
	MaxAssignees      uint      `gorm:"not null"`
	MinAssignees      uint      `gorm:"not null"`
	MaxAttendees      uint      `gorm:"not null"`
	MinAttendees      uint      `gorm:"not null"`
	BookingDeadlineAt time.Time `gorm:"type:datetime;not null"`
	CancelDeadlineAt  time.Time `gorm:"type:datetime;not null"`

	Assignments   []Assignment
	Unassignments []Unassignment
	Bookings      []Booking
	Cancels       []Cancel
	Control       Control
}

func (*Event) TableName() string {
	return "events"
}
