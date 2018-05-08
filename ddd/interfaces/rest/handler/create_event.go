package handler

import (
	"time"

	"github.com/learning-microservice/core/echox"
	"github.com/learning-microservice/event/ddd/application/usecase"
	"github.com/learning-microservice/event/ddd/domain/model/event"
)

/*
curl -XPOST \
  localhost:19000/v1/events \
  -H 'Content-Type: application/json' \
  -d '
  {
    "category": "lesson",
    "contents": ["business", "free trial"],
    "start_at": "2018-02-26T12:00:00+09:00",
    "end_at":   "2018-02-26T12:30:00+09:00",
    "is_privated": true,
    "assignees": ["TUT:1000"]
  }' | jq .
*/
func MakeCreateEventEndpoint(service usecase.CreateEvent) echox.HandlerFunc {
	return func(c *echox.Context) error {
		req := struct {
			Category   event.Category    `json:"category"     binding:"required"`
			Contents   event.Contents    `json:"contents"`
			StartAt    event.StartAt     `json:"start_at"     binding:"required"`
			EndAt      event.EndAt       `json:"end_at"       binding:"required,gtfield=StartAt"`
			IsPrivated bool              `json:"is_privated"`
			Assignees  event.AssigneeIDs `json:"assignees"    binding:"required,min=1"`
		}{}
		if err := c.BindJSON(&req); err != nil {
			return c.BadRequest(err)
		}

		output, err := service.Create(&usecase.CreateEventInput{
			Category:     req.Category,
			Contents:     req.Contents,
			StartAt:      req.StartAt,
			EndAt:        req.EndAt,
			IsPrivated:   req.IsPrivated,
			Assignees:    req.Assignees,
			MinAssignees: 1,
			MaxAssignees: 1,
			MinAttendees: 1,
			MaxAttendees: 1,
			PublishedAt: event.PublishedAt{
				time.Now(),
			},
			BookingDeadlineAt: event.BookingDeadlineAt{
				req.StartAt.Add(-10 * time.Minute),
			},
			CancelDeadlineAt: event.CancelDeadlineAt{
				req.StartAt.Add(-30 * time.Minute),
			},
		})
		if err != nil {
			return c.BadRequest(err)
		}

		return c.OK(&struct {
			EventID    event.ID          `json:"event_id"`
			Category   event.Category    `json:"category"`
			Contents   event.Contents    `json:"contents,omitempty"`
			StartAt    event.StartAt     `json:"start_at"`
			EndAt      event.EndAt       `json:"end_at"`
			IsPrivated bool              `json:"is_privated"`
			Assignees  event.AssigneeIDs `json:"assignees,omitempty"`
			Attendees  event.AttendeeIDs `json:"attendees,omitempty"`
		}{
			EventID:    output.ID(),
			Category:   output.Category(),
			Contents:   output.Contents(),
			StartAt:    output.StartAt(),
			EndAt:      output.EndAt(),
			IsPrivated: output.IsPrivated(),
			Assignees:  output.AssigneeIDs(),
			Attendees:  output.AttendeeIDs(),
		})
	}
}
