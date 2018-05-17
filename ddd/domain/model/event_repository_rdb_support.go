package model

import (
	"encoding/json"

	"github.com/learning-microservice/event/ddd/domain/model/shared/event"
	"github.com/learning-microservice/event/ddd/infrastructure/mysql/records"
)

type EventRepositoryRDBSupport struct {
	AssignmentRepositoryRDBSupport
	BookingRepositoryRDBSupport
}

func (r *EventRepositoryRDBSupport) ToEventEntity(record *records.Event) *Event {
	var (
		tags event.Tags
	)
	if len(record.Tags) > 0 {
		if err := json.Unmarshal(record.Tags, &tags); err != nil {
			// TODO warning log ?
			panic(err)
		}
	}
	return &Event{
		id:         event.ID(record.ID),
		category:   event.Category(record.Category),
		tags:       tags,
		startAt:    event.StartAt{record.StartAt},
		endAt:      event.EndAt{record.EndAt},
		assignment: *r.ToAssignmentEntity(&record.Assignment),
		booking:    *r.ToBookingEntity(&record.Booking),
		version:    record.Version,
	}
}

func (r *EventRepositoryRDBSupport) ToEventRecord(entity *Event) *records.Event {
	return &records.Event{
		ID:         uint(entity.id),
		Category:   string(entity.category),
		Tags:       entity.tags.JSON(),
		StartAt:    entity.startAt.Time,
		EndAt:      entity.endAt.Time,
		Assignment: *r.ToAssignmentRecord(&entity.assignment),
		Booking:    *r.ToBookingRecord(&entity.booking),
		Version:    entity.version,
	}
}
