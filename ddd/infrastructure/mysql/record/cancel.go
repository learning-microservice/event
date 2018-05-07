package record

import (
	"time"
)

type Cancel struct {
	ID         uint64    `gorm:"primary_key"`
	BookingID  uint64    `gorm:"not null"`
	CanceledAt time.Time `gorm:"type:datetime;not null"`
	Reason     string    `gorm:"not null"`
}

func (*Cancel) TableName() string {
	return "event_cancels"
}
