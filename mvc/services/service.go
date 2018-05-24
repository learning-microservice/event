package services

import (
	"context"
	"time"

	"github.com/learning-microservice/event/mvc/commons/db"
	"github.com/learning-microservice/event/mvc/commons/generator"
	"github.com/learning-microservice/event/mvc/commons/types/event"
)

var (
	now = time.Now
)

type Service interface {
	CreateEventService
	UpdateEventService
	FindEventService
	BookEventService
	DeleteEventService
	SearchEventService
}

func NewService() Service {
	return &service{
		withTx:  db.WithTx,
		nextUID: generator.NextUID,
	}
}

type service struct {
	withTx  func(db.TxFunc) error
	nextUID func() event.UID
}

func productCode(_ context.Context) string {
	return "001"
}

func operatorID(_ context.Context) string {
	return "100:1000"
}
