//+build !test

package model

import (
	"testing"
	"time"

	"github.com/learning-microservice/event/ddd/domain/shared/account"
	"github.com/learning-microservice/event/ddd/domain/shared/event"
)

type testingSupport struct {
	*testing.T
}

func (e *testingSupport) Event(
	id event.ID,
	category event.Category,
	tags event.Tags,
	start time.Time,
	end time.Time,
	assigneeID *account.ID,
	attendeeID *account.ID,
	version uint,
) *Event {
	evt := &Event{
		id:         event.ID(id),
		category:   category,
		tags:       tags,
		startAt:    event.StartAt{start},
		endAt:      event.EndAt{end},
		assignment: emptyAssignment,
		booking:    emptyBooking,
		version:    version,
	}
	if assigneeID != nil {
		evt.assignment = Assignment{
			eventID:    evt.id,
			assigneeID: *assigneeID,
		}
	}
	if attendeeID != nil {
		evt.booking = Booking{
			eventID:    evt.id,
			attendeeID: *attendeeID,
		}
	}
	return evt
}

func TestingSupport(t *testing.T) testingSupport {
	return testingSupport{t}
}
