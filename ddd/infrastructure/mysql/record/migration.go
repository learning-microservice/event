package record

import (
	"github.com/learning-microservice/core/db"
)

func Migration() {
	db.Migration(
		&Event{},
		&Assignment{},
		&Unassignment{},
		&Booking{},
		&Cancel{},
		&Control{},
	)
}
