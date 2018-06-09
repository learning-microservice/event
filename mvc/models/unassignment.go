package models

import (
	"time"
)

// Unassignment is ...
type Unassignment struct {
	ID           uint      `gorm:"primary_key"`
	AssignmentID uint      `gorm:"not null"`
	UnassignedAt time.Time `gorm:"type:datetime;not null"`
	Reason       string    `gorm:"not null"`
	OperatorID   string    `gorm:"not null"`
}

// TableName is ...
func (*Unassignment) TableName() string {
	return "event_unassignments"
}
