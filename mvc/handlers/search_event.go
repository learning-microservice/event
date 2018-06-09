package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learning-microservice/event/mvc/services"
)

/*
curl -XGET \
  "localhost:19000/v1/events?category=lesson&tags=business&tags=lesson&start_at=2018-06-26T12:00:00%2B09:00&end_at=2018-06-27T12:00:00%2B09:00&assignee_id=TUT:2000&sort=id:desc&sort=start_at:desc" \
  -H 'Content-Type: application/json' | jq .
*/
func makeSearchEventEndpoint(service services.SearchEventService) gin.HandlerFunc {
	return func(c *gin.Context) {
		input := services.SearchEventInput{}
		if err := c.Bind(&input); err != nil {
			handleError(c, err)
			return
		}
		output, err := service.Search(toContext(c), &input)
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

// SearchEventEndpoint is ...
var SearchEventEndpoint = makeSearchEventEndpoint(service)
