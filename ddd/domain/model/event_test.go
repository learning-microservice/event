package model

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/learning-microservice/event/ddd/domain/shared/account"
	"github.com/learning-microservice/event/ddd/domain/shared/event"
	"github.com/stretchr/testify/assert"
)

func TestAssign(t *testing.T) {
	var (
		start, _ = time.Parse(time.RFC3339, "2018-04-01T10:00:00+09:00")
		end, _   = time.Parse(time.RFC3339, "2018-04-01T10:30:00+09:00")
	)

	type input struct {
		assigneeID account.ID
		event      *Event
	}

	tests := []struct {
		name     string
		input    input
		expected interface{}
	}{
		{
			input: input{
				assigneeID: account.ID("200:2000"),
				event: &Event{
					id:         event.ID(1000),
					category:   event.Category("lesson"),
					tags:       event.Tags{"tag-a", "tag-b"},
					startAt:    event.StartAt{start},
					endAt:      event.EndAt{end},
					assignment: emptyAssignment,
					booking:    emptyBooking,
					version:    9999,
				},
			},
			expected: &Event{
				id:       event.ID(1000),
				category: event.Category("lesson"),
				tags:     event.Tags{"tag-a", "tag-b"},
				startAt:  event.StartAt{start},
				endAt:    event.EndAt{end},
				assignment: Assignment{
					eventID:    event.ID(1000),
					assigneeID: account.ID("200:2000"),
				},
				booking: emptyBooking,
				version: 9999,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var (
				assigneeID = test.input.assigneeID
				evt        = test.input.event
			)
			err := evt.Assign(assigneeID)
			if expectedErr, ok := test.expected.(error); ok {
				assert.Equal(t, expectedErr, err)
			} else {
				assert.Equal(t, test.expected, evt)
			}
		})
	}
}

func TestBook(t *testing.T) {
	var (
		start, _ = time.Parse(time.RFC3339, "2018-04-01T10:00:00+09:00")
		end, _   = time.Parse(time.RFC3339, "2018-04-01T10:30:00+09:00")
	)

	type input struct {
		attendeeID account.ID
		event      *Event
	}

	tests := []struct {
		name     string
		input    input
		expected interface{}
	}{
		{
			input: input{
				attendeeID: account.ID("100:1000"),
				event: &Event{
					id:         event.ID(1000),
					category:   event.Category("lesson"),
					tags:       event.Tags{"tag-a", "tag-b"},
					startAt:    event.StartAt{start},
					endAt:      event.EndAt{end},
					assignment: emptyAssignment,
					booking:    emptyBooking,
					version:    9999,
				},
			},
			expected: &Event{
				id:         event.ID(1000),
				category:   event.Category("lesson"),
				tags:       event.Tags{"tag-a", "tag-b"},
				startAt:    event.StartAt{start},
				endAt:      event.EndAt{end},
				assignment: emptyAssignment,
				booking: Booking{
					eventID:    event.ID(1000),
					attendeeID: account.ID("100:1000"),
				},
				version: 9999,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var (
				attendeeID = test.input.attendeeID
				evt        = test.input.event
			)
			err := evt.Book(attendeeID)
			if expectedErr, ok := test.expected.(error); ok {
				assert.Equal(t, expectedErr, err)
			} else {
				assert.Equal(t, test.expected, evt)
			}
		})
	}
}

func TestNewEvent(t *testing.T) {
	var (
		start, _ = time.Parse(time.RFC3339, "2018-04-01T10:00:00+09:00")
		end, _   = time.Parse(time.RFC3339, "2018-04-01T10:30:00+09:00")

		category = event.Category("lesson")
		tags     = event.Tags{"tag-a", "tag-b"}
		startAt  = event.StartAt{start}
		endAt    = event.EndAt{end}
	)

	actual := NewEvent(
		category,
		tags,
		startAt,
		endAt,
	)

	expected := &Event{
		id:         event.ID(0),
		category:   category,
		tags:       tags,
		startAt:    startAt,
		endAt:      endAt,
		assignment: emptyAssignment,
		booking:    emptyBooking,
		version:    0,
	}

	assert.Equal(t, expected, actual)
	assert.Equal(t, expected.ID(), actual.ID())
	assert.Equal(t, expected.Category(), actual.Category())
	assert.Equal(t, expected.Tags(), actual.Tags())
	assert.Equal(t, expected.StartAt(), actual.StartAt())
	assert.Equal(t, expected.EndAt(), actual.EndAt())
	assert.Equal(t, expected.Assignment(), actual.Assignment())
	assert.Equal(t, expected.Booking(), actual.Booking())
}

func TestMarshalJSON(t *testing.T) {
	var (
		start, _ = time.Parse(time.RFC3339, "2018-04-01T10:00:00+09:00")
		end, _   = time.Parse(time.RFC3339, "2018-04-01T10:30:00+09:00")
	)
	tests := []struct {
		name     string
		input    *Event
		expected string
	}{
		{
			input: &Event{
				id:       event.ID(1000),
				category: event.Category("lesson"),
				tags:     event.Tags{"tag-a", "tag-b"},
				startAt:  event.StartAt{start},
				endAt:    event.EndAt{end},
				assignment: Assignment{
					eventID:    event.ID(1000),
					assigneeID: account.ID("200:2000"),
				},
				booking: Booking{
					eventID:    event.ID(1000),
					attendeeID: account.ID("100:1000"),
				},
				version: 9999,
			},
			expected: `{
				"event_id":   1000,
				"category":   "lesson",
				"tags":       ["tag-a", "tag-b"],
				"start_at":    "2018-04-01T10:00:00+09:00",
				"end_at":      "2018-04-01T10:30:00+09:00",
				"assignee_id": "200:2000",
				"attendee_id": "100:1000"
			}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var (
				actual   map[string]interface{}
				expected map[string]interface{}
			)
			b, _ := json.Marshal(test.input)
			_ = json.Unmarshal(b, &actual)
			_ = json.Unmarshal([]byte(test.expected), &expected)
			assert.Equal(t, expected, actual)
		})
	}
}
