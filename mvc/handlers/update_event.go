package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learning-microservice/event/mvc/commons/types/event"
	"github.com/learning-microservice/event/mvc/services"
)

/*
curl -XPUT \
  localhost:19000/v1/events/xxx \
  -H 'Content-Type: application/json' \
  -d '
  {
    "tags": ["business", "free trial"]
  }' | jq .
*/
func makeUpdateEventEndpoint(service services.UpdateEventService) gin.HandlerFunc {
	return func(c *gin.Context) {
		input := services.UpdateEventInput{
			UID: event.UID(c.Param("id")),
		}
		if err := c.Bind(&input); err != nil {
			handleError(c, err)
			return
		}
		output, err := service.Update(toContext(c), &input)
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

// UpdateEventEndpoint is ...
var UpdateEventEndpoint = makeUpdateEventEndpoint(service)
