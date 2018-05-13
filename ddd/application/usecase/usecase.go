package usecase

import (
	"github.com/learning-microservice/event/ddd/domain/context"
	"github.com/learning-microservice/event/ddd/domain/model/events"
)

type Service interface {
	CreateEvent
}

type service struct {
	context.Transaction
	eventRepos events.Repository
}

func NewService(tx context.Transaction, eventRepos events.Repository) Service {
	return &service{
		Transaction: tx,
		eventRepos:  eventRepos,
	}
}
