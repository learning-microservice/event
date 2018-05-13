package usecase

import (
	"github.com/learning-microservice/event/ddd/domain/context"
	"github.com/learning-microservice/event/ddd/domain/model/events"
)

func (s *service) Assign(id events.ID, assigneeIDs events.AssigneeIDs) (*events.Event, error) {
	var evt *events.Event
	return evt, s.WithTx(func(sess context.Session) (err error) {
		if evt, err = s.eventRepos.FindBy(id)(sess); err != nil {
			return
		}
		if err = s.assign(evt, assigneeIDs)(sess); err != nil {
			return
		}
		// store event
		return s.eventRepos.Store(evt)(sess)
	})
}

////////////////////////////////////////////
// Shared Private Functions
////////////////////////////////////////////

func (s *service) assign(evt *events.Event, assigneeIDs events.AssigneeIDs) func(context.Session) error {
	return func(sess context.Session) (err error) {
		for _, aid := range assigneeIDs {
			// duplicate event for account timeslot
			if s.eventRepos.ExistsBy(
				aid,
				evt.TimeSlot(),
			)(sess) {
				return ErrDuplicateEvent
			}
			// assign event
			if err = evt.Assign(aid); err != nil {
				return
			}
		}
		return
	}
}
