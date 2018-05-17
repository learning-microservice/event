package tests

import (
	"testing"
)

func TestSenario_3(t *testing.T) {
	binding := map[string]interface{}{
		"tutor_3000":  3000,
		"tutor_3001":  3001,
		"student_300": 300,
		"base_date":   baseDate,
	}

	tests := []struct {
		name   string
		input  input
		output output
	}{
		{
			name: "open lesson (tutor_3000)(12:00-12:30)",
			input: input{
				method: "POST",
				url:    "/v1/lessons",
				body: `{
					"tutor_id":  ${tutor_3000},
					"start_at": "${base_date}T12:00:00+09:00",
					"end_at":   "${base_date}T12:30:00+09:00"
				}`,
			},
			output: output{
				status: 201,
				body: `{
					"slot_id": ${slot_id_1}
				}`,
				callback: func(resp map[string]interface{}) {
					if v, ok := resp["slot_id"]; ok {
						binding["slot_id_1"] = v
					}
				},
			},
		},
		{
			name: "open lesson (tutor_3001)(12:00-12:30)",
			input: input{
				method: "POST",
				url:    "/v1/lessons",
				body: `{
					"tutor_id":  ${tutor_3001},
					"start_at": "${base_date}T12:00:00+09:00",
					"end_at":   "${base_date}T12:30:00+09:00"
				}`,
			},
			output: output{
				status: 201,
				body: `{
					"slot_id": ${slot_id_2}
				}`,
				callback: func(resp map[string]interface{}) {
					if v, ok := resp["slot_id"]; ok {
						binding["slot_id_2"] = v
					}
				},
			},
		},
		{
			name: "book lesson (tutor_3000)(12:00-12:30)(student_300)",
			input: input{
				method: "POST",
				url:    "/v1/lessons/${slot_id_1}/students",
				body: `{
					"student_id": ${student_300}
				}`,
			},
			output: output{
				status: 200,
				body: `{
					"slot_id": ${slot_id_1}
				}`,
			},
		},
		{
			name: "transfer lesson (tutor_3000)(12:00-12:30)(student_300) -> (tutor_3001)(12:00-12:30)",
			input: input{
				method: "PUT",
				url:    "/v1/lessons/${slot_id_1}",
				body: `{
					"transfer_slot_id": ${slot_id_2}
				}`,
			},
			output: output{
				status: 200,
				body: `{
					"slot_id": ${slot_id_2}
				}`,
			},
		},
		{
			name: "cancel lesson (tutor_3001)(12:00-12:30)(student_300)",
			input: input{
				method: "DELETE",
				url:    "/v1/lessons/${slot_id_2}/students",
				body: `{
					"student_id": ${student_300}
				}`,
			},
			output: output{
				status: 200,
				body: `{
					"slot_id": ${slot_id_2}
				}`,
			},
		},
		{
			name: "close lesson (tutor_3001)(12:00-12:30)",
			input: input{
				method: "DELETE",
				url:    "/v1/lessons/${slot_id_2}",
			},
			output: output{
				status: 200,
				body: `{
					"slot_id": ${slot_id_2}
				}`,
			},
		},
	}

	runTests(t, tests, &binding)
}
