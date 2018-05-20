package memory

import (
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
	return func(domain.Session) error {
		return nil
	}
}

func (r *assignmentRepository) Delete(entity *model.Assignment, reason string) func(domain.Session) error {
	return func(domain.Session) error {
		return nil
	}
}
