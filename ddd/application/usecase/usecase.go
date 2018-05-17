package usecase

import (
	"github.com/learning-microservice/event/ddd/domain"
	"github.com/learning-microservice/event/ddd/domain/model"
)

type Service interface {
	FindEvent
	CreateEvent
}

type service struct {
	domain.TxContext
	eventRepos  model.EventRepository
	assignRepos model.AssignmentRepository
	bookRepos   model.BookingRepository
}

func NewService(ctx domain.TxContext, repos model.Repositories) Service {
	return &service{
		TxContext:   ctx,
		eventRepos:  repos.EventRepository,
		assignRepos: repos.AssignmentRepository,
		bookRepos:   repos.BookingRepository,
	}
}
