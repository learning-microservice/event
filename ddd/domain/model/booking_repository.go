package model

import (
	"github.com/learning-microservice/event/ddd/domain"
)

type BookingRepository interface {
	Store(book *Booking) func(domain.Session) error
	Delete(book *Booking, reason string) func(domain.Session) error
}
