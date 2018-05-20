package rdb

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
	return func(ses domain.Session) (err error) {
		var (
			mysql       = ses.(*mySQL)
			bookRecords = r.ToBookingRecords(entity)
		)
		for _, record := range bookRecords {
			if mysql.NewRecord(&record) {
				if err = mysql.Create(&record).Error; err != nil {
					return
				}
			}
		}
		return
	}
}

func (r *bookingRepository) Delete(entity *model.Booking, reason string) func(domain.Session) error {
	return func(domain.Session) (err error) {
		return nil
	}
}
