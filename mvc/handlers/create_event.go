package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learning-microservice/event/mvc/services"
)

/*
curl -XPOST \
  localhost:19000/v1/events \
  -H 'Content-Type: application/json' \
  -d '
  {
    "category": "lesson",
    "tags": ["business", "free trial"],
    "start_at": "2018-06-26T12:00:00+09:00",
    "end_at":   "2018-06-26T12:30:00+09:00",
    "assignee_id": "TUT:1000"
  }' | jq .
*/
func makeCreateEventEndpoint(service services.CreateEventService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input services.CreateEventInput
		if err := c.Bind(&input); err != nil {
			handleError(c, err)
			return
		}
		output, err := service.Create(toContext(c), &input)
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusCreated, output)
	}
}

// CreateEventEndpoint is ...
var CreateEventEndpoint = makeCreateEventEndpoint(service)
