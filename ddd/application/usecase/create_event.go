package usecase

import (
	"github.com/learning-microservice/event/ddd/domain"
	"github.com/learning-microservice/event/ddd/domain/model"
	"github.com/learning-microservice/event/ddd/domain/shared/account"
	"github.com/learning-microservice/event/ddd/domain/shared/event"
)

type CreateEvent interface {
	Create(input *CreateEventInput) (*model.Event, error)
}

type CreateEventInput struct {
	Category   event.Category `json:"category"    binding:"required"`
	Tags       event.Tags     `json:"tags"`
	StartAt    event.StartAt  `json:"start_at"    binding:"required"`
	EndAt      event.EndAt    `json:"end_at"      binding:"required,gtfield=StartAt"`
	AssigneeID account.ID     `json:"assignee_id" binding:"required"`
}

func (s *service) Create(input *CreateEventInput) (*model.Event, error) {
	evt := model.NewEvent(
		input.Category,
		input.Tags,
		input.StartAt,
		input.EndAt,
	)
	return evt, s.WithTx(func(ses domain.Session) (err error) {
		// duplicate event for account timeslot
		if s.eventRepos.ExistsBy(
			input.AssigneeID,
			evt.StartAt(),
			evt.EndAt(),
		)(ses) {
			return ErrDuplicateEvent
		}

		// store event
		if err = s.eventRepos.Store(evt)(ses); err != nil {
			return
		}

		// assign event
		return s.assign(evt, input.AssigneeID)(ses)
	})
}
