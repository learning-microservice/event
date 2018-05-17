package mysql

import (
	//"github.com/jinzhu/gorm"
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
	return func(ses domain.Session) error {
		var (
			mysql      = ses.(*mySQL)
			bookRecord = r.ToBookingRecord(entity)
		)
		if mysql.NewRecord(bookRecord) {
			return mysql.Create(bookRecord).Error
		}
		return nil
	}
}

func (r *bookingRepository) Delete(entity *model.Booking, reason string) func(domain.Session) error {
	return func(domain.Session) (err error) {
		return nil
	}
}
