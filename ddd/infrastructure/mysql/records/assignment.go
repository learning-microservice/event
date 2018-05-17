package records

import (
	"time"
)

type Assignment struct {
	ID         uint      `gorm:"primary_key"`
	EventID    uint      `gorm:"not null"`
	AssigneeID string    `gorm:"not null"`
	AssignedAt time.Time `gorm:"type:datetime;not null"`
	IsDeleted  bool      `gorm:"-"`
}

func (*Assignment) TableName() string {
	return "event_assignments"
}
