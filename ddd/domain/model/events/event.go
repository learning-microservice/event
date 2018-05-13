package events

////////////////////////////////////////////
// Event
////////////////////////////////////////////

type Event struct {
	id              ID
	category        Category
	tags            Tags
	timeSlot        TimeSlot
	publishedAt     PublishedAt
	assignments     Assignments
	bookings        Bookings
	maxAssignees    MaxAssignees
	maxAttendees    MaxAttendees
	bookingDeadline BookingDeadline
	cancelDeadline  CancelDeadline
	version         uint
}

func (e *Event) ID() ID {
	return e.id
}

func (e *Event) Category() Category {
	return e.category
}

func (e *Event) Tags() Tags {
	return e.tags
}

func (e *Event) TimeSlot() TimeSlot {
	return e.timeSlot
}

func (e *Event) PublishedAt() PublishedAt {
	return e.publishedAt
}

func (e *Event) MaxAssignees() MaxAssignees {
	return e.maxAssignees
}

func (e *Event) MaxAttendees() MaxAttendees {
	return e.maxAttendees
}

func (e *Event) BookingDeadline() BookingDeadline {
	return e.bookingDeadline
}

func (e *Event) CancelDeadline() CancelDeadline {
	return e.cancelDeadline
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
	var (
	//now = time.Now()
	)
	// exist assignee ?
	if e.assignments.contains(assigneeID) {
		return ErrAlreadyAssigned
	}

	// validate max assignee
	if int(e.maxAssignees) == len(e.assignments) {
		return ErrMaxAssignees
	}

	// validate deadline
	//if e.timeSlot.startAt.Add(e.assignedDeadline * time.Minute).Before(now) {
	//	return ErrAssignedDeadline
	//}

	// add assignee
	e.assignments.add(assigneeID)

	return
}

func (e *Event) Book(attendeeID AttendeeID) (err error) {
	// exist attendee ?
	if e.bookings.contains(attendeeID) {
		return ErrAlreadyBooked
	}

	// validate max attendees
	if int(e.maxAttendees) == len(e.bookings) {
		return
	}

	// validate deadline
	//if e.timeSlot.startAt.Add(e.bookingDeadline * time.Minute).Before(now) {
	//	return ErrBookngDeadline
	//}

	// add attendee
	e.bookings.add(attendeeID)

	return
}

func (e *Event) IsNew() bool {
	return e.version == 0
}

////////////////////////////////////////////
// Public Static Functions
////////////////////////////////////////////

func New(
	category Category,
	tags Tags,
	start StartAt,
	end EndAt,
	publishedAt PublishedAt,
	maxAssignees MaxAssignees,
	maxAttendees MaxAttendees,
	bookingDeadline BookingDeadline,
	cancelDeadline CancelDeadline,
) *Event {
	return &Event{
		category: category,
		tags:     tags,
		timeSlot: TimeSlot{
			startAt: start,
			endAt:   end,
		},
		publishedAt:     publishedAt,
		maxAssignees:    maxAssignees,
		maxAttendees:    maxAttendees,
		bookingDeadline: bookingDeadline,
		cancelDeadline:  cancelDeadline,
	}
}
