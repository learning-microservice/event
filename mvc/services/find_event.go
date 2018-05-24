package services

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/learning-microservice/event/mvc/commons/db"
	"github.com/learning-microservice/event/mvc/commons/errors"
	"github.com/learning-microservice/event/mvc/commons/types/event"
	"github.com/learning-microservice/event/mvc/models"
)

type FindEventService interface {
	Find(context.Context, *FindEventInput) (*models.Event, error)
}

type FindEventInput struct {
	UID event.UID `json:"id" binding:"required"`
}

func (s *service) Find(ctx context.Context, input *FindEventInput) (*models.Event, error) {
	return findBy(input.UID, productCode(ctx))(db.DB())
}

func findBy(uid event.UID, productCode string) func(*gorm.DB) (*models.Event, error) {
	return func(tx *gorm.DB) (*models.Event, error) {
		var evt models.Event
		err := tx.Scopes(
			preloadAssignments,
			preloadBookings,
			filteredUID(uid),
			filteredProductCode(productCode),
		).First(
			&evt,
		).Error

		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.NewNotFoundError(
				"id",
				uid,
				"event not found",
			)
		}
		return &evt, err
	}
}
