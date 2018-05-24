package services

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/learning-microservice/event/mvc/commons/errors"
	"github.com/learning-microservice/event/mvc/commons/types/account"
	"github.com/learning-microservice/event/mvc/commons/types/event"
	"github.com/learning-microservice/event/mvc/models"
)

type CreateEventService interface {
	Create(context.Context, *CreateEventInput) (*models.Event, error)
}

type CreateEventInput struct {
	Category   event.Category `json:"category"    binding:"required"`
	Tags       event.Tags     `json:"tags"`
	StartAt    time.Time      `json:"start_at"    binding:"required,gte"`
	EndAt      time.Time      `json:"end_at"      binding:"required,gtfield=StartAt"`
	AssigneeID account.ID     `json:"assignee_id" binding:"required"`
}

func (s *service) Create(ctx context.Context, input *CreateEventInput) (*models.Event, error) {
	var (
		evt         *models.Event
		cnt         uint
		productCode = productCode(ctx)
		operatorID  = operatorID(ctx)
	)
	return evt, s.withTx(func(tx *gorm.DB) (err error) {
		if cnt, err = countsBy(
			input.AssigneeID,
			input.StartAt,
			input.EndAt,
		)(tx); err != nil {
			return err
		}

		// verify duplicate event for account timeslot
		if cnt > 0 {
			return errors.NewValidationError(
				"assignee_id",
				input.AssigneeID,
				"account already assigned to another event",
			)
		}

		evt = &models.Event{
			UID:         s.nextUID(),
			ProductCode: productCode,
			Category:    input.Category,
			Tags:        input.Tags,
			StartAt:     input.StartAt,
			EndAt:       input.EndAt,
			Version:     1,
			CreatedBy:   operatorID,
			CreatedAt:   now(),
		}

		if err = evt.Assign(&models.Assignment{
			AssigneeID: input.AssigneeID,
			AssignedAt: now(),
			OperatorID: operatorID,
		}); err != nil {
			return err
		}

		return tx.Create(evt).Error
	})
}
