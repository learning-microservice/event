package services

import (
	"context"
	"time"

	"github.com/learning-microservice/event/mvc/commons/logger"
	"github.com/learning-microservice/event/mvc/models"
)

type Logging interface {
	Service
}

type logging struct {
	next *service
}

func (l *logging) Create(c context.Context, input *CreateEventInput) (evt *models.Event, err error) {
	defer func(begin time.Time) {
		logger.Info("create event", logger.Fields{
			"input": input,
			"took":  time.Since(begin).String(),
			"event": evt,
			"err":   err,
		})
	}(time.Now())
	return l.next.Create(c, input)
}

func (l *logging) Update(c context.Context, input *UpdateEventInput) (evt *models.Event, err error) {
	defer func(begin time.Time) {
		logger.Info("update event", logger.Fields{
			"took":  time.Since(begin).String(),
			"event": evt,
			"err":   err,
		})
	}(time.Now())
	return l.next.Update(c, input)
}

func (l *logging) Find(c context.Context, input *FindEventInput) (evt *models.Event, err error) {
	defer func(begin time.Time) {
		logger.Info("find event", logger.Fields{
			"took":  time.Since(begin).String(),
			"event": evt,
			"err":   err,
		})
	}(time.Now())
	return l.next.Find(c, input)
}

func (l *logging) Book(c context.Context, input *BookEventInput) (evt *models.Event, err error) {
	defer func(begin time.Time) {
		logger.Info("book event", logger.Fields{
			"took":  time.Since(begin).String(),
			"event": evt,
			"err":   err,
		})
	}(time.Now())
	return l.next.Book(c, input)
}

func (l *logging) Cancel(c context.Context, input *CancelEventInput) (evt *models.Event, err error) {
	defer func(begin time.Time) {
		logger.Info("cancel event", logger.Fields{
			"took":  time.Since(begin).String(),
			"event": evt,
			"err":   err,
		})
	}(time.Now())
	return l.next.Cancel(c, input)
}

func (l *logging) Delete(c context.Context, input *DeleteEventInput) (evt *models.Event, err error) {
	defer func(begin time.Time) {
		logger.Info("delete event", logger.Fields{
			"took":  time.Since(begin).String(),
			"event": evt,
			"err":   err,
		})
	}(time.Now())
	return l.next.Delete(c, input)
}

func (l *logging) Search(c context.Context, input *SearchEventInput) (evts []*models.Event, err error) {
	defer func(begin time.Time) {
		logger.Info("search event", logger.Fields{
			"took":   time.Since(begin).String(),
			"events": evts,
			"err":    err,
		})
	}(time.Now())
	return l.next.Search(c, input)
}
