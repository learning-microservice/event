package event

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// StartAt is ...
type StartAt struct{ time.Time }

// Value is ...
func (s StartAt) Value() (driver.Value, error) {
	return s.Time, nil
}

// Scan is ...
func (s *StartAt) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), s.Time)
}

// EndAt is ...
type EndAt struct{ time.Time }

// Value is ...
func (e EndAt) Value() (driver.Value, error) {
	return e.Time, nil
}

// TimeSlot is ...
type TimeSlot struct {
	startAt StartAt
	endAt   EndAt
}
