package tests

import (
	"testing"
)

func TestSenario_1(t *testing.T) {
	binding := map[string]interface{}{
		"tutor_1000": 1000,
		"base_date":  baseDate,
	}

	tests := []struct {
		name   string
		input  input
		output output
	}{
		{
			name: "open lesson (tutor_1000)(12:00-12:30)",
			input: input{
				method: "POST",
				url:    "/v1/lessons",
				body: `{
					"tutor_id":  ${tutor_1000},
					"start_at": "${base_date}T12:00:00+09:00",
					"end_at":   "${base_date}T12:30:00+09:00"
				}`,
			},
			output: output{
				status: 201,
				body: `{
					"slot_id": ${slot_id}
				}`,
				callback: func(resp map[string]interface{}) {
					if v, ok := resp["slot_id"]; ok {
						binding["slot_id"] = v
					}
				},
			},
		},
		{
			name: "find lesson (tutor_1000)(12:00-12:30)(status=open)",
			input: input{
				method: "GET",
				url:    "/v1/lessons/${slot_id}",
			},
			output: output{
				status: 200,
				body: `{
					"slot_id":   ${slot_id},
					"tutor_id":  ${tutor_1000},
					"start_at": "${base_date}T12:00:00+09:00",
					"end_at":   "${base_date}T12:30:00+09:00"
				}`,
			},
		},
		{
			name: "close lesson (tutor_1000)(12:00-12:30)",
			input: input{
				method: "DELETE",
				url:    "/v1/lessons/${slot_id}",
			},
			output: output{
				status: 200,
				body: `{
					"slot_id": ${slot_id}
				}`,
			},
		},
		{
			name: "find lesson (tutor_1000)(12:00-12:30)(not found)",
			input: input{
				method: "GET",
				url:    "/v1/lessons/${slot_id}",
			},
			output: output{
				status: 404,
			},
		},
	}

	runTests(t, tests, &binding)
}
