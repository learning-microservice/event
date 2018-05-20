package handler

import (
	"github.com/learning-microservice/core/echox"
	"github.com/learning-microservice/event/ddd/application/usecase"
	"github.com/learning-microservice/event/ddd/domain/shared/event"
)

/*
curl -XGET \
  localhost:19000/v1/events/1 \
  -H 'Content-Type: application/json' | jq .
*/
func MakeFindEventEndpoint(service usecase.FindEvent) echox.HandlerFunc {
	return func(c *echox.Context) error {
		var req = struct {
			usecase.FindEventInput
		}{
			usecase.FindEventInput{
				EventID: event.ID(c.ParamsWithUint("event_id")),
			},
		}
		if err := c.BindJSON(&req); err != nil {
			return c.BadRequest(err)
		}

		output, err := service.Find(
			&req.FindEventInput,
		)
		if err != nil {
			return c.BadRequest(err)
		}
		return c.OK(output)
	}
}
