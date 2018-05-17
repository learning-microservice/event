package mysql

import (
	"sync"
	"testing"
	"time"

	"github.com/learning-microservice/event/ddd/domain"
	"github.com/learning-microservice/event/ddd/domain/model"
	"github.com/learning-microservice/event/ddd/domain/model/shared/account"
	"github.com/learning-microservice/event/ddd/domain/model/shared/event"
	"github.com/learning-microservice/event/ddd/infrastructure/mysql/records"
)

type repositoryMock struct {
	model.EventRepositoryRDBSupport
	mtx   sync.RWMutex
	cache map[uint]records.Event
}

func (r *repositoryMock) FindBy(id event.ID) func(domain.Session) (*model.Event, error) {
	return func(domain.Session) (*model.Event, error) {
		for _, eventRecord := range r.cache {
			if id == event.ID(eventRecord.ID) {
				return r.ToEventEntity(&eventRecord), nil
			}
		}
		return nil, model.ErrNotFound
	}
}

func (r *repositoryMock) ExistsBy(aid account.ID, start event.StartAt, end event.EndAt) func(domain.Session) bool {
	var (
		startTime = start.Time
		endTime   = end.Time
	)
	return func(domain.Session) bool {
		for _, eventRecord := range r.cache {
			var (
				evtStartAt = eventRecord.StartAt
				evtEndAt   = eventRecord.EndAt
			)
			// ({start} <= start_at < {end}) and ({start} < end_at <= {end})
			if (startTime.Equal(evtStartAt) || startTime.Before(evtStartAt)) && evtStartAt.After(endTime) {
				if startTime.Before(evtEndAt) && evtEndAt.After(endTime) {
					if eventRecord.Assignment.AssigneeID == string(aid) {
						return true
					}
					if eventRecord.Booking.AttendeeID == string(aid) {
						return true
					}
				}
			}
		}
		return false
	}
}

func (r *repositoryMock) Store(evt *model.Event) func(domain.Session) error {
	return func(domain.Session) error {
		r.mtx.Lock()
		defer r.mtx.Unlock()

		eventID := uint(evt.ID())
		if eventID == 0 {
			eventID = r.nextID()
		}
		r.cache[eventID] = *r.ToEventRecord(evt)
		return nil
	}
}

func (r *repositoryMock) nextID() uint {
	var maxID uint
	for _, eventRecord := range r.cache {
		if eventRecord.ID > maxID {
			maxID = eventRecord.ID
		}
	}
	return maxID + 1
}

func NewEventRepositoryMock(t *testing.T) model.EventRepository {
	repos := &repositoryMock{
		cache: make(map[uint]records.Event),
	}
	for _, eventRecord := range []records.Event{
		Event_OPEN_ASSIGNEE_2011000_SLOT_1200_1230,
		Event_OPEN_ASSIGNEE_2011000_SLOT_1230_1300,
		Event_OPEN_ASSIGNEE_2011000_SLOT_1300_1330,
	} {
		storeEvent := eventRecord
		repos.cache[storeEvent.ID] = storeEvent
	}
	return repos
}

var (
	BaseTime, _ = time.Parse(time.RFC3339, "2018-04-01T12:00:00+09:00")
	Duration    = 30 * time.Minute
)
var (
	Event_OPEN_ASSIGNEE_2011000_SLOT_1200_1230 = records.Event{
		ID:       1,
		Category: "lesson",
		Tags:     []byte(`["business"]`),
		StartAt:  BaseTime,
		EndAt:    BaseTime.Add(Duration),
		Assignment: records.Assignment{
			ID:         1,
			AssigneeID: "201:1000",
		},
	}
	Event_OPEN_ASSIGNEE_2011000_SLOT_1230_1300 = records.Event{
		ID:       2,
		Category: "lesson",
		Tags:     []byte(`["business"]`),
		StartAt:  BaseTime.Add(1 * Duration),
		EndAt:    BaseTime.Add(2 * Duration),
		Assignment: records.Assignment{
			ID:         1,
			AssigneeID: "201:1000",
		},
	}
	Event_OPEN_ASSIGNEE_2011000_SLOT_1300_1330 = records.Event{
		ID:       2,
		Category: "lesson",
		Tags:     []byte(`["business"]`),
		StartAt:  BaseTime.Add(2 * Duration),
		EndAt:    BaseTime.Add(3 * Duration),
		Assignment: records.Assignment{
			ID:         1,
			AssigneeID: "201:1000",
		},
	}
)
