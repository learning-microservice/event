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
		startAt:     record.StartAt,
		endAt:       record.EndAt,
		isPrivated:  record.IsPrivated,
		publishedAt: record.PublishedAt,
		assignments: assignments,
		bookings:    bookings,
		assignmentRule: AssignmentRule{
			minAssignees: record.MinAssignees,
			maxAssignees: record.MaxAssignees,
		},
		bookingRule: BookingRule{
			minAttendees: record.MinAttendees,
			maxAttendees: record.MaxAttendees,
			deadline:     record.BookingDeadlineAt,
		},
		cancelRule: CancelRule{
			deadline: record.CancelDeadlineAt,
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
		if assign.IsCanceled() {
			unassignments = append(unassignments, record.Unassignment{
				AssignmentID: assign.ID(),
				UnassignedAt: now,
				Reason:       "none",
			})
		} else {
			assignments = append(assignments, record.Assignment{
				ID:         assign.ID(),
				EventID:    eventID,
				AssigneeID: string(assign.AssigneeID()),
				AssignedAt: now,
			})
		}
	}
	for _, book := range evt.bookings {
		if book.IsCanceled() {
			cancels = append(cancels, record.Cancel{
				BookingID:  book.ID(),
				CanceledAt: now,
				Reason:     "none",
			})
		} else {
			bookings = append(bookings, record.Booking{
				ID:         book.ID(),
				EventID:    eventID,
				AttendeeID: string(book.AttendeeID()),
				BookedAt:   now,
			})
		}
	}

	return &record.Event{
		ID:                eventID,
		Category:          string(evt.category),
		Contents:          evt.contents.JSON(),
		StartAt:           evt.startAt,
		EndAt:             evt.endAt,
		IsPrivated:        evt.isPrivated,
		PublishedAt:       evt.publishedAt,
		MaxAssignees:      evt.assignmentRule.maxAssignees,
		MinAssignees:      evt.assignmentRule.minAssignees,
		MaxAttendees:      evt.bookingRule.maxAttendees,
		MinAttendees:      evt.bookingRule.minAttendees,
		BookingDeadlineAt: evt.bookingRule.deadline,
		CancelDeadlineAt:  evt.cancelRule.deadline,
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
