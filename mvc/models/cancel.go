package models

import (
	"time"
)

type Cancel struct {
	ID         uint      `gorm:"primary_key"`
	BookingID  uint      `gorm:"not null"`
	CanceledAt time.Time `gorm:"type:datetime;not null"`
	Reason     string    `gorm:"not null"`
	OperatorID string    `gorm:"not null"`
}

func (*Cancel) TableName() string {
	return "event_cancels"
}
