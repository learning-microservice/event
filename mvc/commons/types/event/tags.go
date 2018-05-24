package event

import (
	"database/sql/driver"
	"encoding/json"
)

type Tags []string

func (t *Tags) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), t)
}

func (t Tags) Value() (driver.Value, error) {
	if len(t) > 0 {
		return json.Marshal(t)
	}
	return []byte("[]"), nil
}
