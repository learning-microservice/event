package event

import (
	"database/sql/driver"
)

type UID string

func (u UID) Value() (driver.Value, error) {
	return string(u), nil
}

func (u *UID) IsNotEmpty() bool {
	return len(*u) > 0
}
