package usecase

import (
	"github.com/learning-microservice/event/ddd/domain/context"
	"github.com/learning-microservice/event/ddd/domain/model/event"
)

type Service interface {
	CreateEvent
}

type service struct {
	context.Transaction
	eventRepos event.Repository
}

func NewService(tx context.Transaction, eventRepos event.Repository) Service {
	return &service{
		Transaction: tx,
		eventRepos:  eventRepos,
	}
}
