package events

import (
	"encoding/json"
	"time"

	"github.com/learning-microservice/event/ddd/infrastructure/mysql/records"
)

type Converter struct{}

func (*Converter) ConvertToEntity(record *records.Event) *Event {
	var (
		tags        Tags
		assignments Assignments
		bookings    Bookings
	)
	if len(record.Tags) > 0 {
		if err := json.Unmarshal(record.Tags, &tags); err != nil {
			// TODO warning log ?
			panic(err)
		}
	}
	for _, assign := range record.Assignments {
		assignments = append(assignments, Assignment{
			id:         assign.ID,
			assigneeID: AssigneeID(assign.AssigneeID),
		})
	}
	for _, book := range record.Bookings {
		bookings = append(bookings, Booking{
			id:         book.ID,
			attendeeID: AttendeeID(book.AttendeeID),
		})
	}
	return &Event{
		id:       ID(record.ID),
		category: Category(record.Category),
		tags:     tags,
		timeSlot: TimeSlot{
			startAt: StartAt{record.StartAt},
			endAt:   EndAt{record.EndAt},
		},
		publishedAt:     PublishedAt{record.PublishedAt},
		assignments:     assignments,
		bookings:        bookings,
		maxAssignees:    MaxAssignees(record.MaxAssignees),
		maxAttendees:    MaxAttendees(record.MaxAttendees),
		bookingDeadline: BookingDeadline(record.BookingDeadline),
		cancelDeadline:  CancelDeadline(record.CancelDeadline),
		version:         record.Version,
	}
}

func (*Converter) ConvertToRecord(evt *Event) *records.Event {
	var (
		now           = time.Now()
		eventID       = uint(evt.id)
		assignments   []records.Assignment
		unassignments []records.Unassignment
		bookings      []records.Booking
		cancels       []records.Cancel
	)
	for _, assign := range evt.assignments {
		if assign.isCanceled {
			unassignments = append(unassignments, records.Unassignment{
				AssignmentID: assign.id,
				UnassignedAt: now,
				Reason:       "none",
			})
		} else {
			assignments = append(assignments, records.Assignment{
				ID:         assign.id,
				EventID:    eventID,
				AssigneeID: string(assign.assigneeID),
				AssignedAt: now,
			})
		}
	}
	for _, book := range evt.bookings {
		if book.isCanceled {
			cancels = append(cancels, records.Cancel{
				BookingID:  book.id,
				CanceledAt: now,
				Reason:     "none",
			})
		} else {
			bookings = append(bookings, records.Booking{
				ID:         book.id,
				EventID:    eventID,
				AttendeeID: string(book.attendeeID),
				BookedAt:   now,
			})
		}
	}

	return &records.Event{
		ID:              eventID,
		Category:        string(evt.category),
		Tags:            evt.tags.JSON(),
		StartAt:         evt.timeSlot.startAt.Time,
		EndAt:           evt.timeSlot.endAt.Time,
		PublishedAt:     evt.publishedAt.Time,
		MaxAssignees:    uint(evt.maxAssignees),
		MaxAttendees:    uint(evt.maxAttendees),
		BookingDeadline: uint(evt.bookingDeadline),
		CancelDeadline:  uint(evt.cancelDeadline),
		Version:         evt.version,

		Assignments:   assignments,
		Unassignments: unassignments,
		Bookings:      bookings,
		Cancels:       cancels,
	}
}
