package memory

import (
	"testing"

	"github.com/learning-microservice/event/ddd/domain/model"
	"github.com/learning-microservice/event/ddd/domain/shared/event"
)

func NewRepositories(t *testing.T, data ...*model.Event) model.Repositories {
	eventRepos := &eventRepository{
		cache: make(map[event.ID]*model.Event),
	}
	for _, e := range data {
		eventRepos.Store(e)(nil)
	}
	return model.Repositories{
		EventRepository:      eventRepos,
		AssignmentRepository: &assignmentRepository{},
		BookingRepository:    &bookingRepository{},
	}
}
