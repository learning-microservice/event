package events

import (
	"github.com/learning-microservice/event/ddd/domain/context"
)

type Repository interface {
	FindBy(ID) func(context.Session) (*Event, error)
	ExistsBy(AccountID, TimeSlot) func(context.Session) bool
	Store(*Event) func(context.Session) error
}
