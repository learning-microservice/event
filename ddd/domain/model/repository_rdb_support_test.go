package model

import (
	"testing"
	"time"

	"github.com/learning-microservice/event/ddd/domain/shared/account"
	"github.com/learning-microservice/event/ddd/domain/shared/event"
	"github.com/learning-microservice/event/ddd/infrastructure/rdb/records"
	"github.com/stretchr/testify/assert"
)

func TestToEventEntity(t *testing.T) {
	var (
		currentTime = time.Now()
		start       = currentTime.Add(60 * time.Hour)
		end         = start.Add(60 * time.Minute)
		mockRepos   = struct {
			EventRepositoryRDBSupport
		}{}
	)

	// override application current time
	now = func() time.Time { return currentTime }

	tests := []struct {
		name     string
		input    *records.Event
		expected *Event
	}{
		{
			input: &records.Event{
				ID:          1000,
				Category:    "lesson",
				Tags:        []byte(`["tag-a","tag-b"]`),
				StartAt:     start,
				EndAt:       end,
				Assignments: nil,
				Bookings:    nil,
				Control: records.Control{
					EventID: 1000,
					Version: 9999,
				},
			},
			expected: &Event{
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
		{
			input: &records.Event{
				ID:       1001,
				Category: "lesson",
				Tags:     []byte(`["tag-a","tag-b"]`),
				StartAt:  start,
				EndAt:    end,
				Assignments: records.Assignments{
					{
						EventID:    1001,
						AssigneeID: "200:2000",
						AssignedAt: currentTime,
					},
				},
				Bookings: records.Bookings{
					{
						EventID:    1001,
						AttendeeID: "100:1000",
						BookedAt:   currentTime,
					},
				},
				Control: records.Control{
					EventID: 1001,
					Version: 9999,
				},
			},
			expected: &Event{
				id:       event.ID(1001),
				category: event.Category("lesson"),
				tags:     event.Tags{"tag-a", "tag-b"},
				startAt:  event.StartAt{start},
				endAt:    event.EndAt{end},
				assignment: Assignment{
					eventID:    event.ID(1001),
					assigneeID: account.ID("200:2000"),
				},
				booking: Booking{
					eventID:    event.ID(1001),
					attendeeID: account.ID("100:1000"),
				},
				version: 9999,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := mockRepos.ToEventEntity(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestToEventRecord(t *testing.T) {
	var (
		currentTime = time.Now()
		start       = currentTime.Add(60 * time.Hour)
		end         = start.Add(60 * time.Minute)
		mockRepos   = struct {
			EventRepositoryRDBSupport
		}{}
	)

	// override application current time
	now = func() time.Time { return currentTime }

	tests := []struct {
		name     string
		input    *Event
		expected *records.Event
	}{
		{
			input: &Event{
				id:         event.ID(1000),
				category:   event.Category("lesson"),
				tags:       event.Tags{"tag-a", "tag-b"},
				startAt:    event.StartAt{start},
				endAt:      event.EndAt{end},
				assignment: emptyAssignment,
				booking:    emptyBooking,
				version:    9999,
			},
			expected: &records.Event{
				ID:          1000,
				Category:    "lesson",
				Tags:        []byte(`["tag-a","tag-b"]`),
				StartAt:     start,
				EndAt:       end,
				Assignments: nil,
				Bookings:    nil,
				Control: records.Control{
					EventID: 1000,
					Version: 9999,
				},
			},
		},
		{
			input: &Event{
				id:       event.ID(1001),
				category: event.Category("lesson"),
				tags:     event.Tags{"tag-a", "tag-b"},
				startAt:  event.StartAt{start},
				endAt:    event.EndAt{end},
				assignment: Assignment{
					eventID:    event.ID(1001),
					assigneeID: account.ID("200:2000"),
				},
				booking: Booking{
					eventID:    event.ID(1001),
					attendeeID: account.ID("100:1000"),
				},
				version: 9999,
			},
			expected: &records.Event{
				ID:       1001,
				Category: "lesson",
				Tags:     []byte(`["tag-a","tag-b"]`),
				StartAt:  start,
				EndAt:    end,
				Assignments: records.Assignments{
					{
						EventID:    1001,
						AssigneeID: "200:2000",
						AssignedAt: currentTime,
					},
				},
				Bookings: records.Bookings{
					{
						EventID:    1001,
						AttendeeID: "100:1000",
						BookedAt:   currentTime,
					},
				},
				Control: records.Control{
					EventID: 1001,
					Version: 9999,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := mockRepos.ToEventRecord(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}
