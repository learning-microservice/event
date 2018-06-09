package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/learning-microservice/event/mvc/commons/middlewares"
	"github.com/learning-microservice/event/mvc/commons/validator"
	"github.com/learning-microservice/event/mvc/handlers"
	"github.com/tylerb/graceful"
)

var (
	name     = "event api server"
	version  = "unknown"
	revision = "unknown"
	host, _  = os.Hostname()
	begin    = time.Now()
)

func main() {
	// environments
	os.Setenv("DB_ADDRESS", "root:password@/localdb?charset=utf8&parseTime=True&loc=Local")
	os.Setenv("DB_DEBUG", "true")
	os.Setenv("DB_AUTO_MIGRATE", "true")

	// setup validator
	binding.Validator = validator.NewStructValidator()

	// setup gin engine
	engine := gin.New()
	{
		engine.Use(
			middlewares.Logger(),
			middlewares.Recovery(),
		)
	}

	// root path api
	engine.GET("/", health)

	// setup router
	v1 := engine.Group("/v1")
	{
		// Search Event
		v1.GET("/events", handlers.SearchEventEndpoint)

		// Create Event
		v1.POST("/events", handlers.CreateEventEndpoint)

		// Update Event
		v1.PUT("/events/:id", handlers.UpdateEventEndpoint)

		// Find Event
		v1.GET("/events/:id", handlers.FindEventEndpoint)

		// Delete Event
		v1.DELETE("/events/:id", handlers.DeleteEventEndpoint)

		// Book Event
		v1.POST("/events/:id/booking", handlers.BookEventEndpoint)

		// Cancel Event
		v1.DELETE("/events/:id/booking", handlers.CancelEventEndpoint)
	}

	// run server
	graceful.ListenAndServe(&http.Server{
		Addr:           ":19000",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}, 10*time.Second)
}

func health(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		&struct {
			Name     string `json:"name"`
			Version  string `json:"version"`
			Revision string `json:"revision"`
			Host     string `json:"host"`
			Started  string `json:"started"`
			Running  string `json:"running"`
			Status   string `json:"status"`
		}{
			Name:     name,
			Version:  version,
			Revision: revision,
			Host:     host,
			Started:  begin.Format(time.RFC3339),
			Running: fmt.Sprintf("%v",
				time.Since(begin).Round(time.Millisecond),
			),
			Status: "ok",
		},
	)
}
