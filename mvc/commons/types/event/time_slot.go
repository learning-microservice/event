package event

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type StartAt struct{ time.Time }

func (s StartAt) Value() (driver.Value, error) {
	return s.Time, nil
}

func (s *StartAt) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), s.Time)
}

type EndAt struct{ time.Time }

func (e EndAt) Value() (driver.Value, error) {
	return e.Time, nil
}

type TimeSlot struct {
	startAt StartAt
	endAt   EndAt
}
