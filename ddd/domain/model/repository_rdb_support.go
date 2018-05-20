package model

import (
	"time"

	"github.com/learning-microservice/event/ddd/domain/shared/account"
	"github.com/learning-microservice/event/ddd/domain/shared/event"
	"github.com/learning-microservice/event/ddd/infrastructure/rdb/records"
)

////////////////////////////////////////////
// EventRepositoryRDBSupport
////////////////////////////////////////////

type EventRepositoryRDBSupport struct {
	AssignmentRepositoryRDBSupport
	BookingRepositoryRDBSupport
}

func (r *EventRepositoryRDBSupport) ToEventEntity(record *records.Event) *Event {
	return &Event{
		id:         event.ID(record.ID),
		category:   event.Category(record.Category),
		tags:       event.ToTags(record.Tags),
		startAt:    event.StartAt{record.StartAt},
		endAt:      event.EndAt{record.EndAt},
		assignment: r.ToAssignmentEntity(record.Assignments),
		booking:    r.ToBookingEntity(record.Bookings),
		version:    record.Control.Version,
	}
}

func (r *EventRepositoryRDBSupport) ToEventRecord(entity *Event) *records.Event {
	return &records.Event{
		ID:          uint(entity.id),
		Category:    string(entity.category),
		Tags:        entity.tags.JSON(),
		StartAt:     entity.startAt.Time,
		EndAt:       entity.endAt.Time,
		Assignments: r.ToAssignmentRecords(&entity.assignment),
		Bookings:    r.ToBookingRecords(&entity.booking),
		Control: records.Control{
			EventID: uint(entity.id),
			Version: entity.version,
		},
	}
}

////////////////////////////////////////////
// AssignmentRepositoryRDBSupport
////////////////////////////////////////////

type AssignmentRepositoryRDBSupport struct{}

func (ar *AssignmentRepositoryRDBSupport) ToAssignmentEntity(list records.Assignments) Assignment {
	if size := len(list); size == 0 {
		return emptyAssignment
	} else if size > 1 {
		// warning log ???
	}
	return Assignment{
		eventID:    event.ID(list[0].EventID),
		assigneeID: account.ID(list[0].AssigneeID),
	}
}

func (*AssignmentRepositoryRDBSupport) ToAssignmentRecords(entity *Assignment) (list records.Assignments) {
	if *entity != emptyAssignment {
		list = append(list, records.Assignment{
			EventID:    uint(entity.eventID),
			AssigneeID: string(entity.assigneeID),
			AssignedAt: now(),
		})
	}
	return
}

////////////////////////////////////////////
// BookingRepositoryRDBSupport
////////////////////////////////////////////

type BookingRepositoryRDBSupport struct{}

func (*BookingRepositoryRDBSupport) ToBookingEntity(list records.Bookings) Booking {
	if size := len(list); size == 0 {
		return emptyBooking
	} else if size > 1 {
		// warning log ???
	}
	return Booking{
		eventID:    event.ID(list[0].EventID),
		attendeeID: account.ID(list[0].AttendeeID),
	}
}

func (*BookingRepositoryRDBSupport) ToBookingRecords(entity *Booking) (list records.Bookings) {
	if *entity != emptyBooking {
		list = append(list, records.Booking{
			EventID:    uint(entity.eventID),
			AttendeeID: string(entity.attendeeID),
			BookedAt:   now(),
		})
	}
	return
}

var now = func() time.Time {
	return time.Now()
}
