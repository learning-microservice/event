package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalJSON(t *testing.T) {
	tests := []struct {
		input    applicationError
		expected string
	}{
		{
			input:    newApplicationError("id", "", "validation error", nil),
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			actual, err := test.input.MarshalJSON()
			if assert.NoError(t, err) {
				assert.Equal(t, test.expected, string(actual))
			}
		})
	}
}
