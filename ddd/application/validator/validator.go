package validator

import (
	"github.com/learning-microservice/core/validator"
)

var (
	validate = validator.NewStructValidator()
)

func ValidateStruct(obj interface{}) error {
	return validate.ValidateStruct()
}
