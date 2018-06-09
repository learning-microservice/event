package event

import (
	"database/sql/driver"
	"encoding/json"
)

// Tags is ...
type Tags []string

// IsEmpty is ...
func (t Tags) IsEmpty() bool {
	return len(t) == 0
}

// IsNotEmpty is ...
func (t Tags) IsNotEmpty() bool {
	return !t.IsEmpty()
}

// Scan is ...
func (t *Tags) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), t)
}

// Value is ...
func (t Tags) Value() (driver.Value, error) {
	if len(t) > 0 {
		return json.Marshal(t)
	}
	return []byte("[]"), nil
}
