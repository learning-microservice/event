package mysql

import (
	"github.com/learning-microservice/event/ddd/domain/model"
)

func NewRepositories() model.Repositories {
	return model.Repositories{
		EventRepository:      &eventRepository{},
		AssignmentRepository: &assignmentRepository{},
		BookingRepository:    &bookingRepository{},
	}
}
