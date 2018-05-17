package tests

import (
	"testing"
)

func TestSenario_4(t *testing.T) {
	binding := map[string]interface{}{
		"tutor_id_1":   4000,
		"tutor_id_2":   4001,
		"student_id_1": 400,
		"student_id_2": 401,
		"base_date":    baseDate,
	}

	tests := []struct {
		name   string
		input  input
		output output
	}{
		// open lesson (tutor_1)(12:00-12:30, 12:30-13:00, 13:00-13:30)
		{
			name: "open lesson (tutor_1)(12:00-12:30)",
			input: input{
				method: "POST",
				url:    "/v1/lessons",
				body: `{
					"tutor_id":  ${tutor_id_1},
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
			name: "open lesson (tutor_1)(12:30-13:00)",
			input: input{
				method: "POST",
				url:    "/v1/lessons",
				body: `{
					"tutor_id":  ${tutor_id_1},
					"start_at": "${base_date}T12:30:00+09:00",
					"end_at":   "${base_date}T13:00:00+09:00"
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
			name: "open lesson (tutor_1)(13:00-13:30)",
			input: input{
				method: "POST",
				url:    "/v1/lessons",
				body: `{
					"tutor_id":  ${tutor_id_1},
					"start_at": "${base_date}T13:00:00+09:00",
					"end_at":   "${base_date}T13:30:00+09:00"
				}`,
			},
			output: output{
				status: 201,
				body: `{
					"slot_id": ${slot_id_3}
				}`,
				callback: func(resp map[string]interface{}) {
					if v, ok := resp["slot_id"]; ok {
						binding["slot_id_3"] = v
					}
				},
			},
		},

		// open lesson (tutor_2)(12:00-12:30)
		{
			name: "open lesson (tutor_2)(12:00-12:30)",
			input: input{
				method: "POST",
				url:    "/v1/lessons",
				body: `{
					"tutor_id":  ${tutor_id_2},
					"start_at": "${base_date}T12:00:00+09:00",
					"end_at":   "${base_date}T12:30:00+09:00"
				}`,
			},
			output: output{
				status: 201,
				body: `{
					"slot_id": ${slot_id_4}
				}`,
				callback: func(resp map[string]interface{}) {
					if v, ok := resp["slot_id"]; ok {
						binding["slot_id_4"] = v
					}
				},
			},
		},

		// book lesson (tutor_1)(12:00-12:30)(student_id_1)
		{
			name: "book lesson (tutor_1)(12:00-12:30)(student_id_1)",
			input: input{
				method: "POST",
				url:    "/v1/lessons/${slot_id_1}/students",
				body: `{
					"student_id": ${student_id_1}
				}`,
			},
			output: output{
				status: 200,
				body: `{
					"slot_id": ${slot_id_1}
				}`,
			},
		},

		// search lesson (tutor_1)
		{
			name: "search lesson (tutor_1)(11:30-12:00)(hits=0)",
			input: input{
				method: "GET",
				url:    "/v1/lessons",
				body: `{
					"tutor_id":  ${tutor_id_1},
					"start_at": "${base_date}T11:30:00+09:00",
					"end_at":   "${base_date}T12:00:00+09:00"
				}`,
			},
			output: output{
				status: 200,
				body:   `[]`,
			},
		},
		{
			name: "search lesson (tutor_1)(12:00-13:30)(hits=3)",
			input: input{
				method: "GET",
				url:    "/v1/lessons",
				body: `{
					"tutor_id":  ${tutor_id_1},
					"start_at": "${base_date}T12:00:00+09:00",
					"end_at":   "${base_date}T13:30:00+09:00"
				}`,
			},
			output: output{
				status: 200,
				body: `[
					{
						"slot_id":   ${slot_id_3},
						"tutor_id":  ${tutor_id_1},
						"start_at": "${base_date}T13:00:00+09:00",
						"end_at":   "${base_date}T13:30:00+09:00"
					},
					{
						"slot_id":   ${slot_id_2},
						"tutor_id":  ${tutor_id_1},
						"start_at": "${base_date}T12:30:00+09:00",
						"end_at":   "${base_date}T13:00:00+09:00"
					},
					{
						"slot_id":    ${slot_id_1},
						"tutor_id":   ${tutor_id_1},
						"student_id": ${student_id_1},
						"start_at":  "${base_date}T12:00:00+09:00",
						"end_at":    "${base_date}T12:30:00+09:00"
					}
				]`,
			},
		},
		{
			name: "search lesson (tutor_1)(13:30-14:00)(hits=0)",
			input: input{
				method: "GET",
				url:    "/v1/lessons",
				body: `{
					"tutor_id":  ${tutor_id_1},
					"start_at": "${base_date}T13:30:00+09:00",
					"end_at":   "${base_date}T14:00:00+09:00"
				}`,
			},
			output: output{
				status: 200,
				body:   `[]`,
			},
		},

		// search lesson (student_id_1)
		{
			name: "search lesson (student_id_1)(11:30-12:00)(hits=0)",
			input: input{
				method: "GET",
				url:    "/v1/lessons",
				body: `{
					"student_id":  ${student_id_1},
					"start_at":   "${base_date}T11:30:00+09:00",
					"end_at":     "${base_date}T12:00:00+09:00"
				}`,
			},
			output: output{
				status: 200,
				body:   `[]`,
			},
		},
		{
			name: "search lesson (student_id_1)(12:00-13:30)(hits=1)",
			input: input{
				method: "GET",
				url:    "/v1/lessons",
				body: `{
					"student_id":  ${student_id_1},
					"start_at":   "${base_date}T12:00:00+09:00",
					"end_at":     "${base_date}T13:30:00+09:00"
				}`,
			},
			output: output{
				status: 200,
				body: `[
					{
						"slot_id":    ${slot_id_1},
						"tutor_id":   ${tutor_id_1},
						"student_id": ${student_id_1},
						"start_at":  "${base_date}T12:00:00+09:00",
						"end_at":    "${base_date}T12:30:00+09:00"
					}
				]`,
			},
		},
		{
			name: "search lesson (student_id_1)(13:30-14:00)(hits=0)",
			input: input{
				method: "GET",
				url:    "/v1/lessons",
				body: `{
					"student_id":  ${student_id_1},
					"start_at":   "${base_date}T13:30:00+09:00",
					"end_at":     "${base_date}T14:00:00+09:00"
				}`,
			},
			output: output{
				status: 200,
				body:   `[]`,
			},
		},
	}

	runTests(t, tests, &binding)
}
