package tests

import (
	"testing"
)

func TestSenario_2(t *testing.T) {
	binding := map[string]interface{}{
		"tutor_2000":  2000,
		"student_200": 200,
		"base_date":   baseDate,
	}

	tests := []struct {
		name   string
		input  input
		output output
	}{
		{
			name: "open lesson (tutor_2000)(12:00-12:30)",
			input: input{
				method: "POST",
				url:    "/v1/lessons",
				body: `{
					"tutor_id": ${tutor_2000},
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
			name: "book lesson (tutor_2000)(12:00-12:30)(student_200)",
			input: input{
				method: "POST",
				url:    "/v1/lessons/${slot_id}/students",
				body: `{
					"student_id": ${student_200}
				}`,
			},
			output: output{
				status: 200,
				body: `{
					"slot_id": ${slot_id}
				}`,
			},
		},
		{
			name: "find lesson (tutor_2000)(12:00-12:30)(student_200)(status=booked)",
			input: input{
				method: "GET",
				url:    "/v1/lessons/${slot_id}",
			},
			output: output{
				status: 200,
				body: `{
					"slot_id":    ${slot_id},
					"tutor_id":   ${tutor_2000},
					"student_id": ${student_200},
					"start_at":  "${base_date}T12:00:00+09:00",
					"end_at":    "${base_date}T12:30:00+09:00"
				}`,
			},
		},
		{
			name: "cancel lesson (tutor_2000)(12:00-12:30)(student_200)",
			input: input{
				method: "DELETE",
				url:    "/v1/lessons/${slot_id}/students",
				body: `{
					"student_id": ${student_200}
				}`,
			},
			output: output{
				status: 200,
				body: `{
					"slot_id": ${slot_id}
				}`,
			},
		},
		{
			name: "find lesson (tutor_2000)(12:00-12:30)(status=open)",
			input: input{
				method: "GET",
				url:    "/v1/lessons/${slot_id}",
			},
			output: output{
				status: 200,
				body: `{
					"slot_id":   ${slot_id},
					"tutor_id":  ${tutor_2000},
					"start_at": "${base_date}T12:00:00+09:00",
					"end_at":   "${base_date}T12:30:00+09:00"
				}`,
			},
		},
		{
			name: "close lesson (tutor_2000)(12:00-12:30)",
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
	}

	runTests(t, tests, &binding)
}
