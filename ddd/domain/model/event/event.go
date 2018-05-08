package event

import (
	"encoding/json"
	"time"
)

////////////////////////////////////////////
// Event
////////////////////////////////////////////

type Event struct {
	id             ID
	category       Category
	contents       Contents
	startAt        StartAt
	endAt          EndAt
	isPrivated     bool
	publishedAt    PublishedAt
	assignments    Assignments
	bookings       Bookings
	assignmentRule AssignmentRule
	bookingRule    BookingRule
	cancelRule     CancelRule
	version        uint
}

func (e *Event) ID() ID {
	return e.id
}

func (e *Event) Category() Category {
	return e.category
}

func (e *Event) Contents() Contents {
	return e.contents
}

func (e *Event) StartAt() StartAt {
	return e.startAt
}

func (e *Event) EndAt() EndAt {
	return e.endAt
}

func (e *Event) IsPrivated() bool {
	return e.isPrivated
}

func (e *Event) PublishedAt() PublishedAt {
	return e.publishedAt
}

func (e *Event) AssigneeIDs() AssigneeIDs {
	var assigneeIDs AssigneeIDs
	for _, assign := range e.assignments {
		if !assign.isCanceled {
			assigneeIDs = append(assigneeIDs, assign.assigneeID)
		}
	}
	return assigneeIDs
}

func (e *Event) AttendeeIDs() AttendeeIDs {
	var attendeeIDs AttendeeIDs
	for _, book := range e.bookings {
		if !book.isCanceled {
			attendeeIDs = append(attendeeIDs, book.attendeeID)
		}
	}
	return attendeeIDs
}

func (e *Event) Assign(assigneeID AssigneeID) (err error) {
	// add assignee
	if err = e.assignments.add(e.id, assigneeID); err != nil {
		return
	}

	// assign rule check
	return //e.bookingRule.validate(len(e.bookings))
}

func (e *Event) Book(attendeeID AttendeeID) (err error) {
	// add attendee
	if err = e.bookings.add(e.id, attendeeID); err != nil {
		return
	}

	// booking rule check
	return e.bookingRule.validate(len(e.bookings))
}

////////////////////////////////////////////
// Value Objects
////////////////////////////////////////////

type (
	ID          string
	Category    string
	Content     string
	Contents    []Content
	StartAt     struct{ time.Time }
	EndAt       struct{ time.Time }
	AssigneeID  string
	AssigneeIDs []AssigneeID
	AttendeeID  string
	AttendeeIDs []AttendeeID

	PublishedAt       struct{ time.Time }
	MaxAssignees      uint
	MinAssignees      uint
	MaxAttendees      uint
	MinAttendees      uint
	BookingDeadlineAt struct{ time.Time }
	CancelDeadlineAt  struct{ time.Time }
)

func (c Contents) JSON() []byte {
	if len(c) > 0 {
		b, _ := json.Marshal(c)
		return b
	} else {
		return []byte("[]")
	}
}

////////////////////////////////////////////
// Public Static Functions
////////////////////////////////////////////

func New(
	category Category,
	contents Contents,
	start StartAt,
	end EndAt,
	isPrivated bool,
	minAssignees MinAssignees,
	maxAssignees MaxAssignees,
	minAttendees MinAttendees,
	maxAttendees MaxAttendees,
	publishedAt PublishedAt,
	bookingDeadlineAt BookingDeadlineAt,
	cancelDeadlineAt CancelDeadlineAt,
) *Event {
	return &Event{
		category:    category,
		contents:    contents,
		startAt:     start,
		endAt:       end,
		isPrivated:  isPrivated,
		publishedAt: publishedAt,
		assignmentRule: AssignmentRule{
			minAssignees: minAssignees,
			maxAssignees: maxAssignees,
		},
		bookingRule: BookingRule{
			minAttendees: minAttendees,
			maxAttendees: maxAttendees,
			deadline:     bookingDeadlineAt,
		},
		cancelRule: CancelRule{
			deadline: cancelDeadlineAt,
		},
		version: 1,
	}
}
