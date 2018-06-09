package event

import (
	"database/sql/driver"
)

// UID is ...
type UID string

// IsEmpty is ...
func (u UID) IsEmpty() bool {
	return len(u) == 0
}

// IsNotEmpty is ...
func (u UID) IsNotEmpty() bool {
	return !u.IsEmpty()
}

// Value is ...
func (u UID) Value() (driver.Value, error) {
	return string(u), nil
}
