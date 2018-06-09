package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learning-microservice/event/mvc/commons/types/event"
	"github.com/learning-microservice/event/mvc/services"
)

/*
curl -XDELETE \
  localhost:19000/v1/events/xxxx/booking \
  -H 'Content-Type: application/json' \
  -d '
  {
    "attendee_id": "STU:1000"
  }' | jq .
*/
func makeCancelEventEndpoint(service services.CancelEventService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input = services.CancelEventInput{
			UID: event.UID(c.Param("id")),
		}
		if err := c.Bind(&input); err != nil {
			handleError(c, err)
			return
		}
		output, err := service.Cancel(toContext(c), &input)
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

// CancelEventEndpoint is ...
var CancelEventEndpoint = makeCancelEventEndpoint(service)
