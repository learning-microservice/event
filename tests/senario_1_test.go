package tests

import (
	"testing"
)

// TestSenario 1
// - Create Event
// - Find Event
// - Book Event
// - Find Event
// - Cancel Event
// - Find Event
// - Delete Event
// - Find Event
func TestSenario_1(t *testing.T) {
	binding := map[string]interface{}{
		"tutor_1000":   "2000:1000",
		"student_1000": "1000:1000",
		"base_date":    baseDate,
	}

	tests := []struct {
		name   string
		input  input
		output output
	}{
		{
			name: "create event (tutor_1000)(12:00-12:30)",
			input: input{
				method: "POST",
				url:    "/v1/events",
				body: `{
					"category":    "lesson",
					"tags":        ["tag-1"],
					"start_at":    "${base_date}T12:00:00+09:00",
					"end_at":      "${base_date}T12:30:00+09:00",
					"assignee_id": "${tutor_1000}"
				}`,
			},
			output: output{
				status: 201,
				body: `{
					"id":          "${id}",
					"category":    "lesson",
					"tags":        ["tag-1"],
					"start_at":    "${base_date}T12:00:00+09:00",
					"end_at":      "${base_date}T12:30:00+09:00",
					"assignee_id": "${tutor_1000}"
				}`,
				callback: func(resp map[string]interface{}) {
					if v, ok := resp["id"]; ok {
						binding["id"] = v
					}
				},
			},
		},
		{
			name: "find event (tutor_1000)(12:00-12:30)(tag-1)(status=opened)",
			input: input{
				method: "GET",
				url:    "/v1/events/${id}",
			},
			output: output{
				status: 200,
				body: `{
					"id":          "${id}",
					"category":    "lesson",
					"tags":        ["tag-1"],
					"start_at":    "${base_date}T12:00:00+09:00",
					"end_at":      "${base_date}T12:30:00+09:00",
					"assignee_id": "${tutor_1000}"
				}`,
			},
		},
		{
			name: "update event (tutor_1000)(12:00-12:30)(tag-1,tag-2)",
			input: input{
				method: "PUT",
				url:    "/v1/events/${id}",
				body: `{
					"tags":        ["tag-1", "tag-2"]
				}`,
			},
			output: output{
				status: 200,
				body: `{
					"id":          "${id}",
					"category":    "lesson",
					"tags":        ["tag-1", "tag-2"],
					"start_at":    "${base_date}T12:00:00+09:00",
					"end_at":      "${base_date}T12:30:00+09:00",
					"assignee_id": "${tutor_1000}"
				}`,
				callback: func(resp map[string]interface{}) {
					if v, ok := resp["id"]; ok {
						binding["id"] = v
					}
				},
			},
		},
		{
			name: "find event (tutor_1000)(12:00-12:30)(tag-1,tag-2)(status=opened)",
			input: input{
				method: "GET",
				url:    "/v1/events/${id}",
			},
			output: output{
				status: 200,
				body: `{
					"id":          "${id}",
					"category":    "lesson",
					"tags":        ["tag-1", "tag-2"],
					"start_at":    "${base_date}T12:00:00+09:00",
					"end_at":      "${base_date}T12:30:00+09:00",
					"assignee_id": "${tutor_1000}"
				}`,
			},
		},
		{
			name: "book event (tutor_1000)(12:00-12:30)(tag-1,tag-2)(student_1000)",
			input: input{
				method: "POST",
				url:    "/v1/events/${id}/booking",
				body: `{
					"attendee_id": "${student_1000}"
				}`,
			},
			output: output{
				status: 200,
				body: `{
					"id":          "${id}",
					"category":    "lesson",
					"tags":        ["tag-1", "tag-2"],
					"start_at":    "${base_date}T12:00:00+09:00",
					"end_at":      "${base_date}T12:30:00+09:00",
					"assignee_id": "${tutor_1000}",
					"attendee_id": "${student_1000}"
				}`,
			},
		},
		{
			name: "find event (tutor_1000)(12:00-12:30)(tag-1,tag-2)(student_1000)(status=closed)",
			input: input{
				method: "GET",
				url:    "/v1/events/${id}",
			},
			output: output{
				status: 200,
				body: `{
					"id":          "${id}",
					"category":    "lesson",
					"tags":        ["tag-1", "tag-2"],
					"start_at":    "${base_date}T12:00:00+09:00",
					"end_at":      "${base_date}T12:30:00+09:00",
					"assignee_id": "${tutor_1000}",
					"attendee_id": "${student_1000}"
				}`,
			},
		},
		{
			name: "cancel event (tutor_1000)(12:00-12:30)(tag-1,tag-2)(student_1000)",
			input: input{
				method: "DELETE",
				url:    "/v1/events/${id}/booking",
				body: `{
					"attendee_id": "${student_1000}"
				}`,
			},
			output: output{
				status: 200,
				body: `{
					"id":          "${id}",
					"category":    "lesson",
					"tags":        ["tag-1", "tag-2"],
					"start_at":    "${base_date}T12:00:00+09:00",
					"end_at":      "${base_date}T12:30:00+09:00",
					"assignee_id": "${tutor_1000}"
				}`,
			},
		},
		{
			name: "find event (tutor_1000)(12:00-12:30)(tag-1,tag-2)(status=opened)",
			input: input{
				method: "GET",
				url:    "/v1/events/${id}",
			},
			output: output{
				status: 200,
				body: `{
					"id":          "${id}",
					"category":    "lesson",
					"tags":        ["tag-1", "tag-2"],
					"start_at":    "${base_date}T12:00:00+09:00",
					"end_at":      "${base_date}T12:30:00+09:00",
					"assignee_id": "${tutor_1000}"
				}`,
			},
		},
		{
			name: "delete event (tutor_1000)(12:00-12:30)(tag-1,tag-2)",
			input: input{
				method: "DELETE",
				url:    "/v1/events/${id}",
			},
			output: output{
				status: 200,
				body: `{
					"id":          "${id}",
					"category":    "lesson",
					"tags":        ["tag-1", "tag-2"],
					"start_at":    "${base_date}T12:00:00+09:00",
					"end_at":      "${base_date}T12:30:00+09:00",
					"assignee_id": "${tutor_1000}"
				}`,
			},
		},
		{
			name: "find event (tutor_1000)(12:00-12:30)(tag-1,tag-2)(not found)",
			input: input{
				method: "GET",
				url:    "/v1/event/${id}",
			},
			output: output{
				status: 404,
			},
		},
	}

	runTests(t, tests, &binding)
}
