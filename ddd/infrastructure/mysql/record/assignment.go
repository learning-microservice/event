package record

import (
	"time"
)

type Assignment struct {
	ID         uint64    `gorm:"primary_key"`
	EventID    string    `gorm:"not null"`
	AssigneeID string    `gorm:"not null"`
	AssignedAt time.Time `gorm:"type:datetime;not null"`
}

func (*Assignment) TableName() string {
	return "event_assignments"
}
