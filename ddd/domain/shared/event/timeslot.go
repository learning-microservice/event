package event

import (
	"time"
)

type StartAt struct{ time.Time }

type EndAt struct{ time.Time }

type TimeSlot struct {
	startAt StartAt
	endAt   EndAt
}

func (t *TimeSlot) StartAt() StartAt {
	return t.startAt
}

func (t *TimeSlot) EndAt() EndAt {
	return t.endAt
}
