package rdb

import (
	"github.com/jinzhu/gorm"
	"github.com/learning-microservice/event/ddd/domain"
	"github.com/learning-microservice/event/ddd/domain/model"
	"github.com/learning-microservice/event/ddd/domain/shared/account"
	"github.com/learning-microservice/event/ddd/domain/shared/event"
	"github.com/learning-microservice/event/ddd/infrastructure/rdb/records"
)

////////////////////////////////////////////
// eventRepositoryOnMySQL
////////////////////////////////////////////

type eventRepository struct {
	model.EventRepositoryRDBSupport
}

func (r *eventRepository) FindBy(id event.ID) func(domain.Session) (*model.Event, error) {
	return func(ses domain.Session) (*model.Event, error) {
		var (
			mysql = ses.(*mySQL)
			evt   = &records.Event{
				ID: uint(id),
			}
		)
		if err := mysql.Preload(
			"Assignment", func(db *gorm.DB) *gorm.DB {
				return db.Table("event_assignment_view")
			},
		).Preload(
			"Booking", func(db *gorm.DB) *gorm.DB {
				return db.Table("event_booking_view")
			},
		).Find(
			evt,
		).Error; err != nil {
			return nil, err
		}
		return r.ToEventEntity(evt), nil
	}
}

func (r *eventRepository) ExistsBy(aid account.ID, start event.StartAt, end event.EndAt) func(domain.Session) bool {
	return func(sess domain.Session) bool {
		var (
			mysql = sess.(*mySQL)
		)
		return !mysql.Where(
			"(start_at >= ? AND start_at < ?) OR (end_at > ? AND end_at <= ?)",
			start.Time, end.Time, start.Time, end.Time,
		).Where(
			"EXISTS ("+
				"SELECT 'e' FROM event_assignment_view ea WHERE ea.assignee_id = ? AND ea.event_id = events.id "+
				"UNION ALL "+
				"SELECT 'e' FROM event_booking_view eb WHERE eb.attendee_id = ? AND eb.event_id = events.id"+
				")",
			string(aid),
			string(aid),
		).Take(
			&records.Event{},
		).RecordNotFound()
	}
}

func (r *eventRepository) Store(entity *model.Event) func(domain.Session) error {
	return func(ses domain.Session) (err error) {
		var (
			mysql       = ses.(*mySQL)
			eventRecord = r.ToEventRecord(entity)
		)
		if mysql.NewRecord(eventRecord) {
			if err = mysql.Create(eventRecord).Error; err != nil {
				return
			}
		}
		*entity = *r.ToEventEntity(eventRecord)
		return
	}
}

func (r *eventRepository) Delete(entity *model.Event) func(domain.Session) error {
	return func(domain.Session) (err error) {
		// TODO delete event_bookins or event_cancels
		// TODO delete event_assignments or event_unassignments
		// TODO delete event_controls
		// TODO delete event
		return nil
	}
}

func (db *mySQL) verifyVersion(eventRecord *records.Event) (err error) {
	var (
		control     = eventRecord.Control
		nextVersion = control.Version + 1
	)
	if control.Version == 0 {
		// create new event control
		if err = db.Create(&control).Error; err != nil {
			return
		}
	} else {
		ret := db.Model(
			&control,
		).Select(
			"version",
		).Where(
			"version = ?", control.Version,
		).Update(
			"version", nextVersion,
		)
		if err = ret.Error; err != nil {
			return
		}
		if ret.RowsAffected == 0 {
			return model.ErrAlreadyModified
		}
	}
	eventRecord.Control.Version = nextVersion
	return
}
