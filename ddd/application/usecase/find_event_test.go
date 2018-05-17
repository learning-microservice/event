package usecase

import (
	"testing"
	"time"

	"github.com/learning-microservice/event/ddd/domain/context"
	"github.com/learning-microservice/event/ddd/domain/model/events"
	"github.com/learning-microservice/event/ddd/infrastructure/mysql"
	"github.com/stretchr/testify/assert"
)

var (
	now = time.Now()
)

func TestFindEvent(t *testing.T) {
	var (
		baseTime = time.Now()
	)

	service := NewService(
		&contextMock{},
		mysql.NewEventRepositoryMock(t, baseTime),
	)

	tests := []struct {
		name     string
		input    FindEventInput
		expected events.Event
	}{
		{
			input: FindEventInput{
				EventID: events.ID(1),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := service.Find(&test.input)

			assert.Nil(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}

type contextMock struct{}

func (c *contextMock) WithTx(f func(context.Session) error) error {
	return f(c)
}

func (c *contextMock) WithReadOnly(f func(context.Session) error) error {
	return f(c)
}
