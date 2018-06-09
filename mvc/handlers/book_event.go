package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learning-microservice/event/mvc/commons/types/event"
	"github.com/learning-microservice/event/mvc/services"
)

/*
curl -XPOST \
  localhost:19000/v1/events/xxxx/booking \
  -H 'Content-Type: application/json' \
  -d '
  {
    "attendee_id": "STU:1000"
  }' | jq .
*/
func makeBookEventEndpoint(service services.BookEventService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input = services.BookEventInput{
			UID: event.UID(c.Param("id")),
		}
		if err := c.Bind(&input); err != nil {
			handleError(c, err)
			return
		}
		output, err := service.Book(toContext(c), &input)
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

// BookEventEndpoint is ...
var BookEventEndpoint = makeBookEventEndpoint(service)
