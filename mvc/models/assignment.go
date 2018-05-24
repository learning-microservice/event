package models

import (
	"time"

	"github.com/learning-microservice/event/mvc/commons/errors"
	"github.com/learning-microservice/event/mvc/commons/types/account"
	"github.com/learning-microservice/event/mvc/commons/types/event"
)

type Assignments []*Assignment

func (as Assignments) existsAttendee(assigneeID account.ID) bool {
	for _, a := range as {
		if a.AssigneeID == assigneeID {
			return true
		}
	}
	return false
}

type Assignment struct {
	ID         uint       `gorm:"primary_key"`
	EventID    event.ID   `gorm:"not null"`
	AssigneeID account.ID `gorm:"not null"`
	AssignedAt time.Time  `gorm:"type:datetime;not null"`
	OperatorID string     `gorm:"type:varchar(36);not null"`
}

func (*Assignment) TableName() string {
	return "event_assignments"
}

func (e *Event) Assign(assignment *Assignment) error {
	// TODO validate deadline

	// validate already assigned
	if e.Assignments.existsAttendee(assignment.AssigneeID) {
		return errors.NewValidationError(
			"assignee_id",
			assignment.AssigneeID,
			"account already assigned",
		)
	}

	// validate max assignees
	if len(e.Bookings) == MAX_ASSIGNEES {
		return errors.NewValidationError(
			"assignee_id",
			assignment.AssigneeID,
			"maximum number of assignees is reached",
		)
	}

	// set event id
	assignment.EventID = e.ID

	// append assignment
	e.Assignments = append(e.Assignments, assignment)
	return nil
}

const (
	MAX_ASSIGNEES = 1
)
