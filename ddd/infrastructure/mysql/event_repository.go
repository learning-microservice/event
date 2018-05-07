package mysql

import (
	"github.com/learning-microservice/core/db"
	"github.com/learning-microservice/event/ddd/domain/context"
	"github.com/learning-microservice/event/ddd/domain/model/event"
	"github.com/rs/xid"
	//"github.com/learning-microservice/event/ddd/infrastructure/mysql/record"
)

type eventRepositoryOnMySQL struct {
	event.Converter
}

func (r *eventRepositoryOnMySQL) FindBy(id string) func(context.Session) (*event.Event, error) {
	return func(context.Session) (*event.Event, error) {
		return nil, nil
	}
}

func (r *eventRepositoryOnMySQL) Store(evt *event.Event) func(context.Session) error {
	return func(sess context.Session) (err error) {
		var (
			mysql  = sess.(*db.DB)
			record = r.ConvertToRecord(evt)
		)

		// create new event
		if mysql.NewRecord(record) {
			record.ID = xid.New().String()
			if err = mysql.Create(record).Error; err != nil {
				return
			}
		} else {
			// TODO create new event bookings
			// TODO create new event cancellations
			// TODO create new event_assignments
			// TODO create new event_unassignments
			// TODO verify event version
		}
		*evt = *r.ConvertToEntity(record)
		return
	}
}

func (r *eventRepositoryOnMySQL) Delete(evt *event.Event) func(context.Session) error {
	return func(context.Session) (err error) {
		// TODO delete event_bookins or event_cancels
		// TODO delete event_assignments or event_unassignments
		// TODO delete event_controls
		// TODO delete event
		return nil
	}
}

func NewEventRepositoryOnMySQL() event.Repository {
	return &eventRepositoryOnMySQL{}
}
