package event

import (
	"encoding/json"
	"time"

	"github.com/learning-microservice/event/ddd/infrastructure/mysql/record"
)

type Converter struct{}

func (*Converter) ConvertToEntity(record *record.Event) *Event {
	var (
		contents    Contents
		assignments Assignments
		bookings    Bookings
	)
	if len(record.Contents) > 0 {
		if err := json.Unmarshal(record.Contents, &contents); err != nil {
			// TODO warning log ?
			panic(err)
		}
	}
	for _, assign := range record.Assignments {
		assignments = append(assignments, Assignment{
			id:         assign.ID,
			eventID:    ID(assign.EventID),
			assigneeID: AssigneeID(assign.AssigneeID),
		})
	}
	for _, book := range record.Bookings {
		bookings = append(bookings, Booking{
			id:         book.ID,
			eventID:    ID(book.EventID),
			attendeeID: AttendeeID(book.AttendeeID),
		})
	}
	return &Event{
		id:          ID(record.ID),
		category:    Category(record.Category),
		contents:    contents,
		startAt:     StartAt{record.StartAt},
		endAt:       EndAt{record.EndAt},
		isPrivated:  record.IsPrivated,
		publishedAt: PublishedAt{record.PublishedAt},
		assignments: assignments,
		bookings:    bookings,
		assignmentRule: AssignmentRule{
			minAssignees: MinAssignees(record.MinAssignees),
			maxAssignees: MaxAssignees(record.MaxAssignees),
		},
		bookingRule: BookingRule{
			minAttendees: MinAttendees(record.MinAttendees),
			maxAttendees: MaxAttendees(record.MaxAttendees),
			deadline:     BookingDeadlineAt{record.BookingDeadlineAt},
		},
		cancelRule: CancelRule{
			deadline: CancelDeadlineAt{record.CancelDeadlineAt},
		},
	}
}

func (*Converter) ConvertToRecord(evt *Event) *record.Event {
	var (
		now           = time.Now()
		eventID       = string(evt.id)
		assignments   []record.Assignment
		unassignments []record.Unassignment
		bookings      []record.Booking
		cancels       []record.Cancel
	)
	for _, assign := range evt.assignments {
		if assign.isCanceled {
			unassignments = append(unassignments, record.Unassignment{
				AssignmentID: assign.id,
				UnassignedAt: now,
				Reason:       "none",
			})
		} else {
			assignments = append(assignments, record.Assignment{
				ID:         assign.id,
				EventID:    eventID,
				AssigneeID: string(assign.assigneeID),
				AssignedAt: now,
			})
		}
	}
	for _, book := range evt.bookings {
		if book.isCanceled {
			cancels = append(cancels, record.Cancel{
				BookingID:  book.id,
				CanceledAt: now,
				Reason:     "none",
			})
		} else {
			bookings = append(bookings, record.Booking{
				ID:         book.id,
				EventID:    eventID,
				AttendeeID: string(book.attendeeID),
				BookedAt:   now,
			})
		}
	}

	return &record.Event{
		ID:                eventID,
		Category:          string(evt.category),
		Contents:          evt.contents.JSON(),
		StartAt:           evt.startAt.Time,
		EndAt:             evt.endAt.Time,
		IsPrivated:        evt.isPrivated,
		PublishedAt:       evt.publishedAt.Time,
		MaxAssignees:      uint(evt.assignmentRule.maxAssignees),
		MinAssignees:      uint(evt.assignmentRule.minAssignees),
		MaxAttendees:      uint(evt.bookingRule.maxAttendees),
		MinAttendees:      uint(evt.bookingRule.minAttendees),
		BookingDeadlineAt: evt.bookingRule.deadline.Time,
		CancelDeadlineAt:  evt.cancelRule.deadline.Time,
		Assignments:       assignments,
		Unassignments:     unassignments,
		Bookings:          bookings,
		Cancels:           cancels,
		Control: record.Control{
			EventID: eventID,
			Version: evt.version,
		},
	}
}
