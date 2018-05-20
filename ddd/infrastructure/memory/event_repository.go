package memory

import (
	"sync"

	"github.com/learning-microservice/event/ddd/domain"
	"github.com/learning-microservice/event/ddd/domain/model"
	"github.com/learning-microservice/event/ddd/domain/shared/account"
	"github.com/learning-microservice/event/ddd/domain/shared/event"
)

type eventRepository struct {
	mtx   sync.RWMutex
	cache map[event.ID]*model.Event
}

func (r *eventRepository) FindBy(id event.ID) func(domain.Session) (*model.Event, error) {
	return func(domain.Session) (*model.Event, error) {
		for _, e := range r.cache {
			if id == e.ID() {
				return e, nil
			}
		}
		return nil, model.ErrNotFound
	}
}

func (r *eventRepository) ExistsBy(aid account.ID, start event.StartAt, end event.EndAt) func(domain.Session) bool {
	var (
		startTime = start.Time
		endTime   = end.Time
	)
	return func(domain.Session) bool {
		for _, e := range r.cache {
			var (
				evtStartAt = e.StartAt().Time
				evtEndAt   = e.EndAt().Time
			)
			// ({start} <= start_at < {end}) and ({start} < end_at <= {end})
			if (startTime.Equal(evtStartAt) || startTime.Before(evtStartAt)) && evtStartAt.After(endTime) {
				if startTime.Before(evtEndAt) && evtEndAt.After(endTime) {
					if e.Assignment().AssigneeID() == aid {
						return true
					}
					if e.Booking().AttendeeID() == aid {
						return true
					}
				}
			}
		}
		return false
	}
}

func (r *eventRepository) Store(evt *model.Event) func(domain.Session) error {
	return func(domain.Session) error {
		r.mtx.Lock()
		defer r.mtx.Unlock()

		eventID := evt.ID()
		if eventID == event.ID(0) {
			eventID = r.nextID()
		}
		r.cache[eventID] = evt
		return nil
	}
}

func (r *eventRepository) nextID() event.ID {
	var maxID uint
	for _, e := range r.cache {
		if uint(e.ID()) > maxID {
			maxID = uint(e.ID())
		}
	}
	return event.ID(maxID + 1)
}
