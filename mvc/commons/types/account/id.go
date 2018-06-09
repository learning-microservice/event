package account

import (
	"database/sql/driver"
)

// ID is ...
type ID string

// Value is ...
func (id ID) Value() (driver.Value, error) {
	return string(id), nil
}

// IsEmpty is ...
func (id ID) IsEmpty() bool {
	return string(id) != ""
}

// IsNotEmpty is ...
func (id ID) IsNotEmpty() bool {
	return !id.IsEmpty()
}
