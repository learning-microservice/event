package model

import (
	"time"

	"github.com/learning-microservice/event/ddd/domain/model/shared/account"
	"github.com/learning-microservice/event/ddd/domain/model/shared/event"
	"github.com/learning-microservice/event/ddd/infrastructure/mysql/records"
)

type BookingRepositoryRDBSupport struct{}

func (*BookingRepositoryRDBSupport) ToBookingEntity(record *records.Booking) *Booking {
	return &Booking{
		id:         record.ID,
		eventID:    event.ID(record.EventID),
		attendeeID: account.ID(record.AttendeeID),
	}
}

func (*BookingRepositoryRDBSupport) ToBookingRecord(entity *Booking) *records.Booking {
	return &records.Booking{
		ID:         entity.id,
		EventID:    uint(entity.eventID),
		AttendeeID: string(entity.attendeeID),
		BookedAt:   time.Now(),
	}
}
