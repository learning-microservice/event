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

	Assignments Assignments `gorm:"-"`
	Bookings    Bookings    `gorm:"-"`
	Control     Control     `gorm:"-"`
}

func (*Event) TableName() string {
	return "events"
}
