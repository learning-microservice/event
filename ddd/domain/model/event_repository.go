package model

import (
	"github.com/learning-microservice/event/ddd/domain"
	"github.com/learning-microservice/event/ddd/domain/shared/account"
	"github.com/learning-microservice/event/ddd/domain/shared/event"
)

type EventRepository interface {
	FindBy(event.ID) func(domain.Session) (*Event, error)
	ExistsBy(account.ID, event.StartAt, event.EndAt) func(domain.Session) bool
	Store(*Event) func(domain.Session) error
}
