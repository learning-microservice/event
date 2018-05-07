package record

import (
	"time"
)

type Unassignment struct {
	ID           uint64    `gorm:"primary_key"`
	AssignmentID uint64    `gorm:"not null"`
	UnassignedAt time.Time `gorm:"type:datetime;not null"`
	Reason       string    `gorm:"not null"`
}

func (*Unassignment) TableName() string {
	return "event_unassignments"
}
