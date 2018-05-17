package mysql

import (
	//"github.com/jinzhu/gorm"
	"github.com/learning-microservice/event/ddd/domain"
	"github.com/learning-microservice/event/ddd/domain/model"
)

////////////////////////////////////////////
// assignmentRepository
////////////////////////////////////////////

type assignmentRepository struct {
	model.AssignmentRepositoryRDBSupport
}

func (r *assignmentRepository) Store(entity *model.Assignment) func(domain.Session) error {
	return func(ses domain.Session) error {
		var (
			mysql        = ses.(*mySQL)
			assignRecord = r.ToAssignmentRecord(entity)
		)
		if mysql.NewRecord(assignRecord) {
			return mysql.Create(assignRecord).Error
		}
		return nil
	}
}

func (r *assignmentRepository) Delete(entity *model.Assignment, reason string) func(domain.Session) error {
	return func(domain.Session) (err error) {
		return nil
	}
}
