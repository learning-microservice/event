package services

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/learning-microservice/event/mvc/commons/errors"
	"github.com/learning-microservice/event/mvc/commons/types/event"
	"github.com/learning-microservice/event/mvc/models"
)

type UpdateEventService interface {
	Update(context.Context, *UpdateEventInput) (*models.Event, error)
}

type UpdateEventInput struct {
	UID  event.UID  `json:"_"   binding:"required"`
	Tags event.Tags `json:"tags"`
}

func (s *service) Update(ctx context.Context, input *UpdateEventInput) (*models.Event, error) {
	var (
		evt         *models.Event
		productCode = productCode(ctx)
	)
	return evt, s.withTx(func(tx *gorm.DB) (err error) {
		if evt, err = findBy(input.UID, productCode)(tx); err != nil {
			return
		}

		var updateColumns []string
		if len(input.Tags) > 0 {
			updateColumns = append(updateColumns, "tags")
			evt.Tags = input.Tags
		}

		// setup event version
		updateColumns = append(updateColumns, "version")
		currentVersion := evt.Version
		evt.Version = currentVersion + 1

		// update event
		sql := tx.Model(evt).Where(
			"version=?", currentVersion,
		).Select(
			updateColumns,
		).Update(evt)

		if err = sql.Error; err != nil {
			return
		}

		if sql.RowsAffected == 0 {
			return errors.NewAlreadyModifiedError(
				"id",
				evt.UID,
				"already modified event")
		}
		return
	})
}
