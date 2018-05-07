package usecase

import (
	"time"

	"github.com/learning-microservice/event/ddd/domain/context"
	"github.com/learning-microservice/event/ddd/domain/model/event"
)

type CreateEvent interface {
	Create(input *CreateEventInput) (*event.Event, error)
}

type CreateEventInput struct {
	Category          event.Category    `json:"category"            binding:"required"`
	Contents          event.Contents    `json:"contents"`
	StartAt           time.Time         `json:"start_at"            binding:"required"`
	EndAt             time.Time         `json:"end_at"              binding:"required,gtfield=StartAt"`
	IsPrivated        bool              `json:"is_privated"`
	PublishedAt       time.Time         `json:"published_at"        binding:"omitempty,ltfield=StartAt`
	Assignees         event.AssigneeIDs `json:"assignees"           binding:"required,min=1"`
	MinAssignees      uint              `json:"min_assignees"       binding:"omitempty,min=1"`
	MaxAssignees      uint              `json:"max_assignees"       binding:"omitempty,min=1,gtefield=MinAssignees"`
	MinAttendees      uint              `json:"min_attendees"       binding:"omitempty,min=1"`
	MaxAttendees      uint              `json:"max_attendees"       binding:"omitempty,min=1,gtefield=MinAttendees"`
	BookingDeadlineAt time.Time         `json:"booking_deadline_at" binding:"omitempty,ltfield=StartAt`
	CancelDeadlineAt  time.Time         `json:"cancel_deadline_at"  binding:"omitempty,ltfield=StartAt`
}

func (s *service) Create(input *CreateEventInput) (*event.Event, error) {
	evt := event.New(
		input.Category,
		input.Contents,
		input.StartAt,
		input.EndAt,
		input.IsPrivated,
		input.PublishedAt,
		input.MinAssignees,
		input.MaxAssignees,
		input.MinAttendees,
		input.MaxAttendees,
		input.BookingDeadlineAt,
		input.CancelDeadlineAt,
	)
	return evt, s.WithTx(func(sess context.Session) (err error) {
		// assign event
		if err = evt.AssignAll(input.Assignees); err != nil {
			return
		}

		// store event
		return s.eventRepos.Store(evt)(sess)
	})
}
