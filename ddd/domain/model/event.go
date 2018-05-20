package model

import (
	"encoding/json"
	"time"

	"github.com/learning-microservice/event/ddd/domain/shared/account"
	"github.com/learning-microservice/event/ddd/domain/shared/event"
)

////////////////////////////////////////////
// Event
////////////////////////////////////////////

type Event struct {
	id         event.ID
	category   event.Category
	tags       event.Tags
	startAt    event.StartAt
	endAt      event.EndAt
	assignment Assignment
	booking    Booking
	version    uint
}

func (e *Event) ID() event.ID {
	return e.id
}

func (e *Event) Category() event.Category {
	return e.category
}

func (e *Event) Tags() event.Tags {
	return e.tags
}

func (e *Event) StartAt() event.StartAt {
	return e.startAt
}

func (e *Event) EndAt() event.EndAt {
	return e.endAt
}

func (e *Event) Assignment() Assignment {
	return e.assignment
}

func (e *Event) Booking() Booking {
	return e.booking
}

func (e *Event) Assign(assigneeID account.ID) (err error) {
	e.assignment = Assignment{
		eventID:    e.id,
		assigneeID: assigneeID,
	}
	return
}

func (e *Event) Book(attendeeID account.ID) (err error) {
	e.booking = Booking{
		eventID:    e.id,
		attendeeID: attendeeID,
	}
	return
}

func (e *Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID         event.ID       `json:"event_id"`
		Category   event.Category `json:"category"`
		Tags       event.Tags     `json:"tags,omitempty"`
		StartAt    time.Time      `json:"start_at"`
		EndAt      time.Time      `json:"end_at"`
		AssigneeID account.ID     `json:"assignee_id,omitempty"`
		AttendeeID account.ID     `json:"attendee_id,omitempty"`
	}{
		ID:         e.id,
		Category:   e.category,
		Tags:       e.tags,
		StartAt:    e.startAt.Time,
		EndAt:      e.endAt.Time,
		AssigneeID: e.assignment.AssigneeID(),
		AttendeeID: e.booking.AttendeeID(),
	})
}

////////////////////////////////////////////
// Public Static Functions
////////////////////////////////////////////

func NewEvent(
	category event.Category,
	tags event.Tags,
	start event.StartAt,
	end event.EndAt,
) *Event {
	return &Event{
		category: category,
		tags:     tags,
		startAt:  start,
		endAt:    end,
	}
}
