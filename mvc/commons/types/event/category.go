package event

import (
	"database/sql/driver"
)

type Category string

func (c Category) Value() (driver.Value, error) {
	return string(c), nil
}

func (c *Category) IsNotEmpty() bool {
	return len(string(*c)) > 0
}
