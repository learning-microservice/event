package records

import (
	"time"
)

type Event struct {
	ID              uint      `gorm:"primary_key"`
	Category        string    `gorm:"not null"`
	Tags            []byte    `gorm:"type:json"`
	StartAt         time.Time `gorm:"type:datetime;not null"`
	EndAt           time.Time `gorm:"type:datetime;not null"`
	PublishedAt     time.Time `gorm:"type:datetime;not null"`
	MaxAssignees    uint      `gorm:"not null"`
	MaxAttendees    uint      `gorm:"not null"`
	BookingDeadline uint      `gorm:"not null"`
	CancelDeadline  uint      `gorm:"not null"`
	Version         uint      `gorm:"-"`

	Assignments   []Assignment
	Bookings      []Booking
	Unassignments []Unassignment `gorm:"-"`
	Cancels       []Cancel       `gorm:"-"`
	Control       Control        `gorm:"-"`
}

func (*Event) TableName() string {
	return "events"
}
