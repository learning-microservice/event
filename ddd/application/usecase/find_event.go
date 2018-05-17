package usecase

import (
	"github.com/learning-microservice/event/ddd/domain"
	"github.com/learning-microservice/event/ddd/domain/model"
	"github.com/learning-microservice/event/ddd/domain/model/shared/event"
)

type FindEvent interface {
	Find(input *FindEventInput) (*model.Event, error)
}

type FindEventInput struct {
	EventID event.ID `json:"event_id" binding:"required"`
}

func (s *service) Find(input *FindEventInput) (*model.Event, error) {
	var evt *model.Event
	return evt, s.WithReadOnly(func(ses domain.Session) (err error) {
		evt, err = s.eventRepos.FindBy(input.EventID)(ses)
		return err
	})
}
