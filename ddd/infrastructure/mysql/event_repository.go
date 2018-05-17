package mysql

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/learning-microservice/event/ddd/domain"
	"github.com/learning-microservice/event/ddd/domain/model"
	"github.com/learning-microservice/event/ddd/domain/model/shared/account"
	"github.com/learning-microservice/event/ddd/domain/model/shared/event"
	"github.com/learning-microservice/event/ddd/infrastructure/mysql/records"
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
		)
		eventRecord, err := mysql.findBy(uint(id))
		if err != nil {
			return nil, err
		}
		return r.ToEventEntity(eventRecord), nil
	}
}

func (r *eventRepository) ExistsBy(aid account.ID, start event.StartAt, end event.EndAt) func(domain.Session) bool {
	return func(sess domain.Session) bool {
		var (
			mysql = sess.(*mySQL)
		)
		return mysql.existsBy(string(aid), start.Time, end.Time)
	}
}

func (r *eventRepository) Store(entity *model.Event) func(domain.Session) error {
	return func(ses domain.Session) (err error) {
		var (
			mysql       = ses.(*mySQL)
			eventRecord = r.ToEventRecord(entity)
		)
		if mysql.NewRecord(eventRecord) {
			if err = mysql.createEvent(eventRecord); err != nil {
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

////////////////////////////////////////////
// Private gorm Functions
////////////////////////////////////////////

func (db *mySQL) findBy(id uint) (*records.Event, error) {
	var evt = &records.Event{
		ID: id,
	}
	return evt, db.Preload(
		"Assignment", func(db *gorm.DB) *gorm.DB {
			return db.Table("event_assignment_view")
		},
	).Preload(
		"Booking", func(db *gorm.DB) *gorm.DB {
			return db.Table("event_booking_view")
		},
	).Find(
		evt,
	).Error
}

func (db *mySQL) existsBy(aid string, start, end time.Time) bool {
	return !db.Where(
		"(start_at >= ? AND start_at < ?) OR (end_at > ? AND end_at <= ?)",
		start, end, start, end,
	).Where(
		"EXISTS ("+
			"SELECT 'e' FROM event_assignment_view ea WHERE ea.assignee_id = ? AND ea.event_id = events.id "+
			"UNION ALL "+
			"SELECT 'e' FROM event_booking_view eb WHERE eb.attendee_id = ? AND eb.event_id = events.id"+
			")",
		aid,
		aid,
	).Take(
		&records.Event{},
	).RecordNotFound()
}

func (db *mySQL) createEvent(eventRecord *records.Event) (err error) {
	// create new event
	return db.Create(eventRecord).Error
}

func (db *mySQL) updateEvent(eventRecord *records.Event) (err error) {
	// TODO create new event_assignment
	if db.NewRecord(&eventRecord.Assignment) {
		if err = db.Create(&eventRecord.Assignment).Error; err != nil {
			return
		}
	}
	// TODO create new event bookings
	if db.NewRecord(&eventRecord.Booking) {
		if err = db.Create(&eventRecord.Booking).Error; err != nil {
			return
		}
	}
	return
}

func (db *mySQL) verifyVersion(eventRecord *records.Event) (err error) {
	control := records.Control{
		EventID: eventRecord.ID,
		Version: eventRecord.Version + 1,
	}
	if eventRecord.Version == 0 {
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
			"version = ?", eventRecord.Version,
		).Update(
			&control,
		)
		if err = ret.Error; err != nil {
			return
		}
		if ret.RowsAffected == 0 {
			return model.ErrAlreadyModified
		}
	}
	eventRecord.Version = control.Version
	return
}
