package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learning-microservice/event/mvc/commons/types/event"
	"github.com/learning-microservice/event/mvc/services"
)

/*
curl -XGET \
  localhost:19000/v1/events/1 \
  -H 'Content-Type: application/json' | jq .
*/
func makeFindEventEndpoint(service services.FindEventService) gin.HandlerFunc {
	return func(c *gin.Context) {
		input := services.FindEventInput{
			UID: event.UID(c.Param("id")),
		}
		if err := c.Bind(&input); err != nil {
			handleError(c, err)
			return
		}
		output, err := service.Find(toContext(c), &input)
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

var FindEventEndpoint = makeFindEventEndpoint(service)
