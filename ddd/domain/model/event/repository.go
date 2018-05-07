package event

import (
	"github.com/learning-microservice/event/ddd/domain/context"
)

type Repository interface {
	FindBy(id string) func(context.Session) (*Event, error)
	Store(*Event) func(context.Session) error
	Delete(*Event) func(context.Session) error
}
