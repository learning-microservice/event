package event

import (
	"time"
)

////////////////////////////////////////////
// AssignmentRule
////////////////////////////////////////////

type AssignmentRule struct {
	minAssignees MinAssignees
	maxAssignees MaxAssignees
	deadline     time.Time
}

func (ar *AssignmentRule) validate(assigneeCount int) error {
	// アサイン上限枠チェック
	if uint(ar.maxAssignees) < uint(assigneeCount) {
		return errorBookingFailure("maximum number of assignees is reached")
	}
	// アサイン締切時間チェック
	if ar.deadline.After(time.Now()) {
		return errorBookingFailure("maximum number of assignees is reached")
	}
	return nil
}

////////////////////////////////////////////
// BookingRule
////////////////////////////////////////////

type BookingRule struct {
	minAttendees MinAttendees
	maxAttendees MaxAttendees
	deadline     BookingDeadlineAt
}

func (br *BookingRule) validate(attendeeCount int) error {
	// 予約上限枠チェック
	if uint(br.maxAttendees) < uint(attendeeCount) {
		return errorBookingFailure("maximum number of attendees is reached")
	}
	// 予約締切時間チェック
	if br.deadline.After(time.Now()) {
		return errorBookingFailure("maximum number of attendees is reached")
	}
	return nil
}

////////////////////////////////////////////
// CancelRule
////////////////////////////////////////////

type CancelRule struct {
	deadline CancelDeadlineAt
}

func (cr *CancelRule) validate() error {
	// 予約キャンセル締切時間チェック
	if cr.deadline.After(time.Now()) {
		return errorBookingFailure("maximum number of attendees is reached")
	}
	return nil
}
