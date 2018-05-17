package usecase

import (
	"github.com/learning-microservice/event/ddd/domain"
	"github.com/learning-microservice/event/ddd/domain/model"
	"github.com/learning-microservice/event/ddd/domain/model/shared/account"
	"github.com/learning-microservice/event/ddd/domain/model/shared/event"
)

type AssignEventInput struct {
	EventID    event.ID   `json:"event_id"    binding:"required"`
	AssigneeID account.ID `json:"assignee_id" binding:"required"`
}

func (s *service) Assign(input *AssignEventInput) (*model.Event, error) {
	var evt *model.Event
	return evt, s.WithTx(func(ses domain.Session) (err error) {
		if evt, err = s.eventRepos.FindBy(input.EventID)(ses); err != nil {
			return
		}

		// duplicate event for account timeslot
		if s.eventRepos.ExistsBy(
			input.AssigneeID,
			evt.StartAt(),
			evt.EndAt(),
		)(ses) {
			return ErrDuplicateEvent
		}

		return s.assign(evt, input.AssigneeID)(ses)
	})
}

////////////////////////////////////////////
// Shared Private Functions
////////////////////////////////////////////

func (s *service) assign(evt *model.Event, assigneeID account.ID) func(domain.Session) error {
	return func(ses domain.Session) (err error) {
		// assign event
		if err = evt.Assign(assigneeID); err != nil {
			return
		}

		// store event assignment
		assignement := evt.Assignment()
		return s.assignRepos.Store(&assignement)(ses)
	}
}
