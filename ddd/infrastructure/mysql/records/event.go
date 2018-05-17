package records

import (
	"time"
)

type Event struct {
	ID       uint      `gorm:"primary_key"`
	Category string    `gorm:"not null"`
	Tags     []byte    `gorm:"type:json"`
	StartAt  time.Time `gorm:"type:datetime;not null"`
	EndAt    time.Time `gorm:"type:datetime;not null"`
	Version  uint      `gorm:"-"`

	Assignment Assignment `gorm:"-"`
	Booking    Booking    `gorm:"-"`
	Control    Control    `gorm:"-"`
}

func (*Event) TableName() string {
	return "events"
}
