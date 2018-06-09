package services

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/learning-microservice/event/mvc/commons/types/account"
	"github.com/learning-microservice/event/mvc/commons/types/event"
	"github.com/learning-microservice/event/mvc/models"
)

// CancelEventService is ...
type CancelEventService interface {
	Cancel(context.Context, *CancelEventInput) (*models.Event, error)
}

// CancelEventInput is ...
type CancelEventInput struct {
	UID        event.UID  `json:"_"           binding:"required"`
	AttendeeID account.ID `json:"attendee_id" binding:"required"`
}

// Cancel is ...
func (s *service) Cancel(ctx context.Context, input *CancelEventInput) (*models.Event, error) {
	var (
		evt *models.Event

		attendeeID  = input.AttendeeID
		productCode = productCode(ctx)
		operatorID  = operatorID(ctx)
	)
	return evt, s.withTx(func(tx *gorm.DB) (err error) {
		if evt, err = findBy(input.UID, productCode)(tx); err != nil {
			return
		}

		cancel := &models.Cancel{
			CanceledAt: now(),
			OperatorID: operatorID,
		}
		if err = evt.Cancel(attendeeID, cancel); err != nil {
			return
		}

		if err = tx.Create(cancel).Error; err != nil {
			return
		}

		// TODO verify version (optimistic locking)
		return
	})
}
