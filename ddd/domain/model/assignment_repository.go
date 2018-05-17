package model

import (
	"github.com/learning-microservice/event/ddd/domain"
)

type AssignmentRepository interface {
	Store(assign *Assignment) func(domain.Session) error
	Delete(assign *Assignment, reason string) func(domain.Session) error
}
