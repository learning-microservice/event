package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/learning-microservice/event/mvc/commons/validator"
	"github.com/learning-microservice/event/mvc/handlers"
	"github.com/tylerb/graceful"
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
		engine.Use(gin.Logger(), gin.Recovery())
	}

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
