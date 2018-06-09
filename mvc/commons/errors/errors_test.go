package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplicationError(t *testing.T) {
	err := newApplicationError(
		ValidationErrorType,
		"id",
		"",
		"validation error",
		nil)
	assert.Equal(t, ValidationErrorType, err.Type())
	assert.Equal(t, "validation error", err.Error())
	assert.Equal(t, nil, err.Cause())
}
