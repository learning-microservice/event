package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learning-microservice/event/mvc/commons/errors"
	"github.com/learning-microservice/event/mvc/commons/validator"
	"github.com/learning-microservice/event/mvc/services"
)

var (
	service = services.NewService()
)

func toContext(_ *gin.Context) context.Context {
	return context.Background()
}

type errorResponse struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func handleError(c *gin.Context, err error) {
	if validatorError, ok := err.(*validator.Errors); ok {
		c.JSON(http.StatusBadRequest, errorResponse{
			Message: validatorError.Error(),
			Errors:  validatorError,
		})
		return
	}

	if jsonError, ok := err.(*json.UnmarshalTypeError); ok {
		c.JSON(http.StatusBadRequest, errorResponse{
			Message: "validation error",
			Errors: []error{
				errors.NewValidationError(
					jsonError.Field,
					nil,
					"invalid json type",
				),
			},
		})
		return
	}

	if appError, ok := err.(errors.ApplicationError); ok {
		var statusCode int
		switch appError.Type() {
		case errors.ValidationErrorType:
			statusCode = http.StatusBadRequest
		case errors.NotFoundErrorType:
			statusCode = http.StatusNotFound
		case errors.AlreadyModifiedErrorType:
			statusCode = http.StatusConflict
		default:
			statusCode = http.StatusInternalServerError
		}
		c.JSON(statusCode, errorResponse{
			Message: appError.Error(),
			Errors:  []error{appError},
		})
		return
	}

	c.JSON(http.StatusInternalServerError, &errorResponse{
		Message: err.Error(),
		Errors: gin.H{
			"errors": []map[string]string{
				{
					"message": err.Error(),
				},
			},
		},
	})
}
