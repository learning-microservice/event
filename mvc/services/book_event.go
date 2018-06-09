package services

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/learning-microservice/event/mvc/commons/errors"
	"github.com/learning-microservice/event/mvc/commons/types/account"
	"github.com/learning-microservice/event/mvc/commons/types/event"
	"github.com/learning-microservice/event/mvc/models"
)

// BookEventService is ...
type BookEventService interface {
	Book(context.Context, *BookEventInput) (*models.Event, error)
}

// BookEventInput is ...
type BookEventInput struct {
	UID        event.UID  `json:"_"           binding:"required"`
	AttendeeID account.ID `json:"attendee_id" binding:"required"`
}

// Book is ...
func (s *service) Book(ctx context.Context, input *BookEventInput) (*models.Event, error) {
	var (
		evt     *models.Event
		booking *models.Booking
		cnt     uint

		attendeeID  = input.AttendeeID
		productCode = productCode(ctx)
		operatorID  = operatorID(ctx)
	)
	return evt, s.withTx(func(tx *gorm.DB) (err error) {
		if evt, err = findBy(input.UID, productCode)(tx); err != nil {
			return
		}

		if cnt, err = countsBy(attendeeID, evt.StartAt, evt.EndAt)(tx); err != nil {
			return
		}

		// verify duplicate event for account timeslot
		if cnt > 0 {
			return errors.NewValidationError(
				"attendee_id",
				attendeeID,
				"account already booked to another event",
			)
		}

		booking = &models.Booking{
			AttendeeID: attendeeID,
			BookedAt:   now(),
			OperatorID: operatorID,
		}
		if err = evt.Book(booking); err != nil {
			return
		}

		if err = tx.Create(booking).Error; err != nil {
			return
		}

		// TODO verify version (optimistic locking)
		return
	})
}
