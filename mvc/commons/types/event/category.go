package event

import (
	"database/sql/driver"
)

// Category is ...
type Category string

// Value is ...
func (c Category) Value() (driver.Value, error) {
	return string(c), nil
}

// IsEmpty is ...
func (c Category) IsEmpty() bool {
	return len(string(c)) > 0
}

// IsNotEmpty is ...
func (c Category) IsNotEmpty() bool {
	return !c.IsEmpty()
}
