package model

import (
	"github.com/learning-microservice/event/ddd/domain/shared/account"
	"github.com/learning-microservice/event/ddd/domain/shared/event"
)

type Assignment struct {
	eventID    event.ID
	assigneeID account.ID
}

func (a Assignment) AssigneeID() account.ID {
	return a.assigneeID
}

var (
	emptyAssignment = Assignment{}
)
