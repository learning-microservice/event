package services

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/learning-microservice/event/mvc/commons/errors"
	"github.com/learning-microservice/event/mvc/commons/types/event"
	"github.com/learning-microservice/event/mvc/models"
)

// DeleteEventService is ...
type DeleteEventService interface {
	Delete(context.Context, *DeleteEventInput) (*models.Event, error)
}

// DeleteEventInput is ...
type DeleteEventInput struct {
	UID event.UID `json:"_" binding:"required"`
}

// Delete is ...
func (s *service) Delete(ctx context.Context, input *DeleteEventInput) (*models.Event, error) {
	var (
		evt         *models.Event
		productCode = productCode(ctx)
	)
	return evt, s.withTx(func(tx *gorm.DB) (err error) {
		if evt, err = findBy(
			input.UID,
			productCode,
		)(tx); err != nil {
			return
		}

		// booked event ?
		if len(evt.Bookings) > 0 {
			return errors.NewValidationError(
				"uid",
				evt.UID,
				"booked event can not be deleted",
			)
		}

		// batch delete event_unassignments
		if err = tx.Where(
			"EXISTS ("+
				"SELECT 'e' FROM event_assignments ea "+
				"WHERE ea.id = event_unassignments.assignment_id "+
				"AND   ea.event_id = ?)",
			evt.ID,
		).Delete(
			&models.Unassignment{},
		).Error; err != nil {
			return
		}

		// batch delete event_assignments
		if err = tx.Where(
			"event_id = ?", evt.ID,
		).Delete(
			&models.Assignment{},
		).Error; err != nil {
			return
		}

		// batch delete event_cancels
		if err = tx.Where(
			"EXISTS ("+
				"SELECT 'e' FROM event_bookings eb "+
				"WHERE eb.id = event_cancels.booking_id "+
				"AND   eb.event_id = ?)",
			evt.ID,
		).Delete(
			&models.Cancel{},
		).Error; err != nil {
			return
		}

		// batch delete event_bookings
		if err = tx.Where(
			"event_id = ?", evt.ID,
		).Delete(
			&models.Booking{},
		).Error; err != nil {
			return
		}

		// delete event
		if err = tx.Delete(evt, "version = ?", evt.Version).Error; err != nil {
			return
		}

		// TODO verify version (optimistic locking)
		return
	})
}
