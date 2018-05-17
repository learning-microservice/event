package model

import (
	"github.com/learning-microservice/event/ddd/domain/model/shared/account"
	"github.com/learning-microservice/event/ddd/domain/model/shared/event"
)

type Assignment struct {
	id         uint
	eventID    event.ID
	assigneeID account.ID
}

func (a *Assignment) AssigneeID() account.ID {
	return a.assigneeID
}

func newAssignment(eventID event.ID, assigneeID account.ID) Assignment {
	return Assignment{
		eventID:    eventID,
		assigneeID: assigneeID,
	}
}
