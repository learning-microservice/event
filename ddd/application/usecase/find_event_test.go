package usecase

import (
	"testing"
	"time"

	"github.com/learning-microservice/event/ddd/domain/model"
	"github.com/learning-microservice/event/ddd/domain/shared/account"
	"github.com/learning-microservice/event/ddd/domain/shared/event"
	"github.com/learning-microservice/event/ddd/infrastructure/memory"
	"github.com/stretchr/testify/assert"
)

var (
	now = time.Now()
)

func TestFindEvent(t *testing.T) {
	var (
		s     = model.TestingSupport(t)
		start = time.Now()
		end   = start.Add(30 * time.Minute)

		assignee_200 = account.ID("201:200")
		attendee_100 = account.ID("101:100")
	)

	data := []*model.Event{
		s.Event(event.ID(1), event.Category("lesson"), event.Tags([]string{"tag-a"}), start, end, nil, nil, 1),
		s.Event(event.ID(2), event.Category("lesson"), event.Tags([]string{"tag-b"}), start, end, &assignee_200, nil, 1),
		s.Event(event.ID(3), event.Category("lesson"), event.Tags([]string{"tag-c"}), start, end, &assignee_200, &attendee_100, 1),
	}

	service := NewService(
		memory.NewTxManager(t),
		memory.NewRepositories(t, data...),
	)

	tests := []struct {
		name     string
		input    FindEventInput
		expected *model.Event
	}{
		{
			input: FindEventInput{
				EventID: event.ID(2),
			},
			expected: data[1],
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
