package account

import (
	"database/sql/driver"
)

type ID string

func (id ID) Value() (driver.Value, error) {
	return string(id), nil
}

func (id ID) IsNotEmpty() bool {
	return string(id) != ""
}
