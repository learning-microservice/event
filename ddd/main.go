package main

import (
	"os"
	"time"

	"github.com/learning-microservice/core/echox"
	"github.com/learning-microservice/core/validator"
	"github.com/learning-microservice/event/ddd/application/usecase"
	"github.com/learning-microservice/event/ddd/infrastructure/mysql"
	"github.com/learning-microservice/event/ddd/infrastructure/mysql/record"
	"github.com/learning-microservice/event/ddd/interfaces/rest/handler"
	"github.com/tylerb/graceful"
)

func main() {
	os.Setenv("DB_ADDRESS", "root:password@/localdb?charset=utf8&parseTime=True&loc=Local")
	os.Setenv("DB_DEBUG", "true")

	// Mysql Transaction Context
	txContext := mysql.NewTransactionContext()

	// Migration
	record.Migration()

	// repositories
	eventRepos := mysql.NewEventRepositoryOnMySQL()

	// service
	service := usecase.NewService(txContext, eventRepos)

	// setup echo engine
	engine := echox.New()
	{
		engine.Server.Addr = ":19000"
		engine.Validator = validator.NewStructValidator()
	}

	// setup middlewares
	{
		//routes.Use(middleware.MakeLoggingMiddleware())
		//routes.Use(gin.Recovery())
	}

	// setup router
	v1 := engine.Group("/v1")
	{
		// Create Event
		v1.POST("/events", echox.Ep(handler.MakeCreateEventEndpoint(service)))
	}

	// run server
	graceful.ListenAndServe(engine.Server, 10*time.Second)
}
