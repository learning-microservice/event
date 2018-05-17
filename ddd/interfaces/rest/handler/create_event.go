package handler

import (
	"github.com/learning-microservice/core/echox"
	"github.com/learning-microservice/event/ddd/application/usecase"
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
    "assignee_id": "TUT:1000"
  }' | jq .
*/
func MakeCreateEventEndpoint(service usecase.CreateEvent) echox.HandlerFunc {
	return func(c *echox.Context) error {
		var req = struct {
			usecase.CreateEventInput
		}{}
		if err := c.BindJSON(&req); err != nil {
			return c.BadRequest(err)
		}

		output, err := service.Create(
			&req.CreateEventInput,
		)
		if err != nil {
			return c.BadRequest(err)
		}
		return c.OK(output)
	}
}
