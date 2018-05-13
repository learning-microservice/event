package handler

import (
	"time"

	"github.com/learning-microservice/core/echox"
	"github.com/learning-microservice/event/ddd/application/usecase"
	"github.com/learning-microservice/event/ddd/domain/model/events"
)

/*
curl -XPOST \
  localhost:19000/v1/events \
  -H 'Content-Type: application/json' \
  -d '
  {
    "category": "lesson",
    "tags": ["business", "free trial"],
    "start_at": "2018-02-26T12:00:00+09:00",
    "end_at":   "2018-02-26T12:30:00+09:00",
    "published_at": "2018-02-25T12:00:00+09:00",
    "assignees": ["TUT:1000"],
    "max_assignees": 1,
    "max_attendees": 1,
    "booking_deadline": 30,
    "cancel_deadline":  10
  }' | jq .
*/
func MakeCreateEventEndpoint(service usecase.CreateEvent) echox.HandlerFunc {
	return func(c *echox.Context) error {
		var req = struct {
			usecase.CreateEventInput
		}{
			CreateEventInput: usecase.CreateEventInput{
				PublishedAt: events.PublishedAt{
					time.Now(),
				},
				MaxAssignees:    1,
				MaxAttendees:    1,
				BookingDeadline: 10,
				CancelDeadline:  30,
			},
		}
		if err := c.BindJSON(&req); err != nil {
			return c.BadRequest(err)
		}

		output, err := service.Create(
			&req.CreateEventInput,
		)
		if err != nil {
			return c.BadRequest(err)
		}

		return c.OK(&struct {
			EventID   events.ID          `json:"event_id"`
			Category  events.Category    `json:"category"`
			Tags      events.Tags        `json:"tags,omitempty"`
			StartAt   events.StartAt     `json:"start_at"`
			EndAt     events.EndAt       `json:"end_at"`
			Assignees events.AssigneeIDs `json:"assignees,omitempty"`
			Attendees events.AttendeeIDs `json:"attendees,omitempty"`
		}{
			EventID:   output.ID(),
			Category:  output.Category(),
			Tags:      output.Tags(),
			StartAt:   output.TimeSlot().StartAt(),
			EndAt:     output.TimeSlot().EndAt(),
			Assignees: output.AssigneeIDs(),
			Attendees: output.AttendeeIDs(),
		})
	}
}
