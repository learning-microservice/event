package models

import (
	"encoding/json"
	"time"

	"github.com/learning-microservice/event/mvc/commons/types/account"
	"github.com/learning-microservice/event/mvc/commons/types/event"
)

type Event struct {
	ID          event.ID       `gorm:"primary_key"`
	UID         event.UID      `gorm:"type:varchar(36);not null"`
	ProductCode string         `gorm:"type:char(3);not null"`
	Category    event.Category `gorm:"type:varchar(20);not null"`
	Tags        event.Tags     `gorm:"type:json"`
	StartAt     time.Time      `gorm:"type:datetime;not null"`
	EndAt       time.Time      `gorm:"type:datetime;not null"`
	Version     uint           `gorm:"type:int;not null"`
	CreatedBy   string         `gorm:"type:varchar(36);not null"`
	CreatedAt   time.Time      `gorm:"type:datetime;not null"`

	Assignments Assignments
	Bookings    Bookings
	//Control     Control     `gorm:"-"`
}

func (*Event) TableName() string {
	return "events"
}

func (e *Event) MarshalJSONss() ([]byte, error) {
	var (
		assigneeID account.ID
		attendeeID account.ID
	)
	if len(e.Assignments) > 0 {
		assigneeID = e.Assignments[0].AssigneeID
	}
	if len(e.Bookings) > 0 {
		attendeeID = e.Bookings[0].AttendeeID
	}

	return json.Marshal(&struct {
		UID        event.UID      `json:"id"`
		Category   event.Category `json:"category"`
		Tags       event.Tags     `json:"tags,omitempty"`
		StartAt    string         `json:"start_at"`
		EndAt      string         `json:"end_at"`
		AssigneeID account.ID     `json:"assignee_id"`
		AttendeeID account.ID     `json:"attendee_id,omitempty"`
	}{
		UID:        e.UID,
		Category:   e.Category,
		Tags:       e.Tags,
		StartAt:    e.StartAt.Format(time.RFC3339),
		EndAt:      e.EndAt.Format(time.RFC3339),
		AssigneeID: assigneeID,
		AttendeeID: attendeeID,
	})
}
