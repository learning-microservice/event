package mysql

import (
	"time"

	"github.com/learning-microservice/event/ddd/domain/context"
	"github.com/learning-microservice/event/ddd/domain/model/events"
	"github.com/learning-microservice/event/ddd/infrastructure/mysql/records"
)

////////////////////////////////////////////
// eventRepositoryOnMySQL
////////////////////////////////////////////

type eventRepositoryOnMySQL struct {
	events.Converter
}

func (r *eventRepositoryOnMySQL) FindBy(id events.ID) func(context.Session) (*events.Event, error) {
	return func(context.Session) (*events.Event, error) {
		return nil, nil
	}
}

func (r *eventRepositoryOnMySQL) ExistsBy(aid events.AccountID, slot events.TimeSlot) func(context.Session) bool {
	return func(sess context.Session) bool {
		var (
			mysql = sess.(*mySQL)
		)
		return mysql.existsBy(aid.UID(), slot.StartAt().Time, slot.EndAt().Time)
	}
}

func (r *eventRepositoryOnMySQL) Store(evt *events.Event) func(context.Session) error {
	return func(sess context.Session) (err error) {
		var (
			mysql       = sess.(*mySQL)
			eventRecord = r.ConvertToRecord(evt)
		)
		if evt.IsNew() {
			if err = mysql.createEvent(eventRecord); err != nil {
				return
			}
		} else {
			if err = mysql.updateEvent(eventRecord); err != nil {
				return
			}
		}
		if err = mysql.verifyVersion(eventRecord); err != nil {
			return
		}
		*evt = *r.ConvertToEntity(eventRecord)
		return
	}
}

func (r *eventRepositoryOnMySQL) Delete(evt *events.Event) func(context.Session) error {
	return func(context.Session) (err error) {
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
	// TODO create new event_assignments
	for _, assign := range eventRecord.Assignments {
		if db.NewRecord(&assign) {
			if err = db.Create(&assign).Error; err != nil {
				return
			}
		}
	}
	// TODO create new event_unassignments
	for _, unassign := range eventRecord.Unassignments {
		if db.NewRecord(&unassign) {
			if err = db.Create(&unassign).Error; err != nil {
				return
			}
		}
	}
	// TODO create new event bookings
	for _, book := range eventRecord.Bookings {
		if db.NewRecord(&book) {
			if err = db.Create(&book).Error; err != nil {
				return
			}
		}
	}
	// TODO create new event cancellations
	for _, cancel := range eventRecord.Cancels {
		if db.NewRecord(&cancel) {
			if err = db.Create(&cancel).Error; err != nil {
				return
			}
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
			return events.ErrAlreadyModified
		}
	}
	eventRecord.Version = control.Version
	return
}

////////////////////////////////////////////
// Public Static Functions
////////////////////////////////////////////

func NewEventRepositoryOnMySQL() events.Repository {
	return &eventRepositoryOnMySQL{}
}
