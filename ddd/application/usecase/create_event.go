package usecase

import (
	"github.com/learning-microservice/event/ddd/domain/context"
	"github.com/learning-microservice/event/ddd/domain/model/event"
)

type CreateEvent interface {
	Create(input *CreateEventInput) (*event.Event, error)
}

type CreateEventInput struct {
	Category          event.Category          `json:"category"            binding:"required"`
	Contents          event.Contents          `json:"contents"`
	StartAt           event.StartAt           `json:"start_at"            binding:"required"`
	EndAt             event.EndAt             `json:"end_at"              binding:"required,gtfield=StartAt"`
	IsPrivated        bool                    `json:"is_privated"`
	Assignees         event.AssigneeIDs       `json:"assignees"           binding:"required,min=1"`
	MinAssignees      event.MinAssignees      `json:"min_assignees"       binding:"omitempty,min=1"`
	MaxAssignees      event.MaxAssignees      `json:"max_assignees"       binding:"omitempty,min=1,gtefield=MinAssignees"`
	MinAttendees      event.MinAttendees      `json:"min_attendees"       binding:"omitempty,min=1"`
	MaxAttendees      event.MaxAttendees      `json:"max_attendees"       binding:"omitempty,min=1,gtefield=MinAttendees"`
	PublishedAt       event.PublishedAt       `json:"published_at"        binding:"omitempty,ltfield=StartAt`
	BookingDeadlineAt event.BookingDeadlineAt `json:"booking_deadline_at" binding:"omitempty,ltfield=StartAt`
	CancelDeadlineAt  event.CancelDeadlineAt  `json:"cancel_deadline_at"  binding:"omitempty,ltfield=StartAt`
}

func (s *service) Create(input *CreateEventInput) (*event.Event, error) {
	evt := event.New(
		input.Category,
		input.Contents,
		input.StartAt,
		input.EndAt,
		input.IsPrivated,
		input.MinAssignees,
		input.MaxAssignees,
		input.MinAttendees,
		input.MaxAttendees,
		input.PublishedAt,
		input.BookingDeadlineAt,
		input.CancelDeadlineAt,
	)
	return evt, s.WithTx(func(sess context.Session) (err error) {
		// assign event
		for _, aid := range input.Assignees {
			if err = evt.Assign(aid); err != nil {
				return
			}
		}

		// store event
		return s.eventRepos.Store(evt)(sess)
	})
}
