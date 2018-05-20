package rdb

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
	return func(ses domain.Session) (err error) {
		var (
			mysql         = ses.(*mySQL)
			assignRecords = r.ToAssignmentRecords(entity)
		)
		for _, record := range assignRecords {
			if mysql.NewRecord(&record) {
				if err = mysql.Create(&record).Error; err != nil {
					return
				}
			}
		}
		return
	}
}

func (r *assignmentRepository) Delete(entity *model.Assignment, reason string) func(domain.Session) error {
	return func(domain.Session) (err error) {
		return nil
	}
}
