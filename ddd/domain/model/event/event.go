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
	startAt        time.Time
	endAt          time.Time
	isPrivated     bool
	publishedAt    time.Time
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

func (e *Event) StartAt() time.Time {
	return e.startAt
}

func (e *Event) EndAt() time.Time {
	return e.endAt
}

func (e *Event) IsPrivated() bool {
	return e.isPrivated
}

func (e *Event) PublishedAt() time.Time {
	return e.publishedAt
}

func (e *Event) AssigneeIDs() AssigneeIDs {
	var assigneeIDs AssigneeIDs
	for _, assign := range e.assignments {
		assigneeIDs = append(assigneeIDs, assign.assigneeID)
	}
	return assigneeIDs
}

func (e *Event) AttendeeIDs() AttendeeIDs {
	var attendeeIDs AttendeeIDs
	for _, book := range e.bookings {
		attendeeIDs = append(attendeeIDs, book.attendeeID)
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

func (e *Event) AssignAll(assigneeIDs []AssigneeID) (err error) {
	for _, aid := range assigneeIDs {
		if err = e.Assign(aid); err != nil {
			return
		}
	}
	return
}

func (e *Event) Book(attendeeID AttendeeID) (err error) {
	// add attendee
	if err = e.bookings.add(e.id, attendeeID); err != nil {
		return
	}

	// booking rule check
	return e.bookingRule.validate(len(e.bookings))
}

func (e *Event) BookAll(attendeeIDs []AttendeeID) (err error) {
	for _, aid := range attendeeIDs {
		if err = e.Book(aid); err != nil {
			return
		}
	}
	return
}

////////////////////////////////////////////
// Value Objects
////////////////////////////////////////////

type (
	ID          string
	Category    string
	Content     string
	Contents    []Content
	AssigneeID  string
	AssigneeIDs []AssigneeID
	AttendeeID  string
	AttendeeIDs []AttendeeID
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
	start time.Time,
	end time.Time,
	isPrivated bool,
	publishedAt time.Time,
	maxAssignees uint,
	minAssignees uint,
	maxAttendees uint,
	minAttendees uint,
	bookingDeadlineAt time.Time,
	cancelDeadlineAt time.Time,
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
