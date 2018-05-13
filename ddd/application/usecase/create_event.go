package usecase

import (
	"github.com/learning-microservice/event/ddd/domain/context"
	"github.com/learning-microservice/event/ddd/domain/model/events"
)

type CreateEvent interface {
	Create(input *CreateEventInput) (*events.Event, error)
}

type CreateEventInput struct {
	Category        events.Category        `json:"category"         binding:"required"`
	Tags            events.Tags            `json:"tags"`
	StartAt         events.StartAt         `json:"start_at"         binding:"required"`
	EndAt           events.EndAt           `json:"end_at"           binding:"required,gtfield=StartAt"`
	PublishedAt     events.PublishedAt     `json:"published_at"`
	Assignees       events.AssigneeIDs     `json:"assignees"        binding:"required,min=1"`
	MaxAssignees    events.MaxAssignees    `json:"max_assignees"`
	MaxAttendees    events.MaxAttendees    `json:"max_attendees"`
	BookingDeadline events.BookingDeadline `json:"booking_deadline" binding:"omitempty,min=1`
	CancelDeadline  events.CancelDeadline  `json:"cancel_deadline"  binding:"omitempty,min=1`
}

func (s *service) Create(input *CreateEventInput) (*events.Event, error) {
	evt := events.New(
		input.Category,
		input.Tags,
		input.StartAt,
		input.EndAt,
		input.PublishedAt,
		input.MaxAssignees,
		input.MaxAttendees,
		input.BookingDeadline,
		input.CancelDeadline,
	)
	return evt, s.WithTx(func(sess context.Session) (err error) {
		if err = s.assign(evt, input.Assignees)(sess); err != nil {
			return
		}
		// store event
		return s.eventRepos.Store(evt)(sess)
	})
}
