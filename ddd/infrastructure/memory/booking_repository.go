package memory

import (
	"github.com/learning-microservice/event/ddd/domain"
	"github.com/learning-microservice/event/ddd/domain/model"
)

////////////////////////////////////////////
// bookingRepository
////////////////////////////////////////////

type bookingRepository struct {
	model.BookingRepositoryRDBSupport
}

func (r *bookingRepository) Store(entity *model.Booking) func(domain.Session) error {
	return func(domain.Session) error {
		return nil
	}
}

func (r *bookingRepository) Delete(entity *model.Booking, reason string) func(domain.Session) error {
	return func(domain.Session) error {
		return nil
	}
}
