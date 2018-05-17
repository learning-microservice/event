package model

import (
	"time"

	"github.com/learning-microservice/event/ddd/domain/model/shared/account"
	"github.com/learning-microservice/event/ddd/domain/model/shared/event"
	"github.com/learning-microservice/event/ddd/infrastructure/mysql/records"
)

type AssignmentRepositoryRDBSupport struct{}

func (*AssignmentRepositoryRDBSupport) ToAssignmentEntity(record *records.Assignment) *Assignment {
	return &Assignment{
		id:         record.ID,
		eventID:    event.ID(record.EventID),
		assigneeID: account.ID(record.AssigneeID),
	}
}

func (*AssignmentRepositoryRDBSupport) ToAssignmentRecord(entity *Assignment) *records.Assignment {
	return &records.Assignment{
		ID:         entity.id,
		EventID:    uint(entity.eventID),
		AssigneeID: string(entity.assigneeID),
		AssignedAt: time.Now(),
	}
}
