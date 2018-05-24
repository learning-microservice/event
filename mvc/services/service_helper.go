package services

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/learning-microservice/event/mvc/commons/types/account"
	"github.com/learning-microservice/event/mvc/models"
)

func countsBy(aid account.ID, from, to time.Time) func(*gorm.DB) (uint, error) {
	return func(tx *gorm.DB) (size uint, err error) {
		err = tx.Model(
			&models.Event{},
		).Where(
			"(start_at >= ? AND start_at < ?) OR (end_at > ? AND end_at <= ?)",
			from, to, from, to,
		).Where(
			"EXISTS ("+
				"SELECT 'e' FROM event_assignment_view ea WHERE ea.assignee_id = ? AND ea.event_id = events.id "+
				"UNION ALL "+
				"SELECT 'e' FROM event_booking_view eb WHERE eb.attendee_id = ? AND eb.event_id = events.id"+
				")",
			aid, aid,
		).Count(
			&size,
		).Error
		return
	}
}
