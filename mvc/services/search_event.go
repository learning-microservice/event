package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/learning-microservice/event/mvc/commons/db"
	"github.com/learning-microservice/event/mvc/commons/types/account"
	"github.com/learning-microservice/event/mvc/commons/types/event"
	"github.com/learning-microservice/event/mvc/models"
)

// SearchEventService is ...
type SearchEventService interface {
	Search(context.Context, *SearchEventInput) ([]*models.Event, error)
}

// SearchEventInput is ...
type SearchEventInput struct {
	Category   event.Category `json:"category"    form:"category"`
	Tags       event.Tags     `json:"tags"        form:"tags"`
	StartAt    time.Time      `json:"start_at"    form:"start_at" time_format:"2006-01-02T15:04:05Z07:00"`
	EndAt      time.Time      `json:"end_at"      form:"end_at"   time_format:"2006-01-02T15:04:05Z07:00" binding:"omitempty,gtfield=StartAt"`
	AssigneeID account.ID     `json:"assignee_id" form:"assignee_id"`
	AttendeeID account.ID     `json:"attendee_id" form:"attendee_id"`
	Sort       []string       `json:"sort"        form:"sort"`
}

func (s *SearchEventInput) parseOrders() (orders []string) {
	for _, s := range s.Sort {
		keyValues := strings.Split(s, ":")
		if len(keyValues) == 2 {
			orders = append(orders, fmt.Sprintf("%s %s", keyValues[0], keyValues[1]))
		}
	}
	return
}

// Search is ...
func (s *service) Search(ctx context.Context, input *SearchEventInput) ([]*models.Event, error) {
	var list []*models.Event
	err := db.DB().Scopes(
		preloadAssignments,
		preloadBookings,
		filteredProductCode(productCode(ctx)),
		filteredCategory(input.Category),
		filteredTags(input.Tags),
		filteredTimeSlot(input.StartAt, input.EndAt),
		filteredAssignee(input.AssigneeID),
		filteredAttendee(input.AttendeeID),
		orders(input.parseOrders()),
	).Find(&list).Error
	return list, err
}

func filteredProductCode(productCode string) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if len(productCode) > 0 {
			return tx.Where(
				"product_code = ?", productCode,
			)
		}
		return tx
	}
}

func filteredID(id event.ID) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if id.IsNotEmpty() {
			return tx.Where(
				"id = ?", id,
			)
		}
		return tx
	}
}

func filteredUID(uid event.UID) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if uid.IsNotEmpty() {
			return tx.Where(
				"uid = ?", uid,
			)
		}
		return tx
	}
}

func filteredCategory(category event.Category) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if category.IsNotEmpty() {
			return tx.Where(
				"category = ?", category,
			)
		}
		return tx
	}
}

func filteredTags(tags event.Tags) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if len(tags) > 0 {
			return tx.Where(
				"JSON_CONTAINS(tags, ?)", tags,
			)
		}
		return tx
	}
}

func filteredTimeSlot(start, end time.Time) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if !start.IsZero() && !end.IsZero() {
			return tx.Where(
				"start_at BETWEEN ? AND ? AND end_at BETWEEN ? AND ?",
				start, end, start, end,
			)
		}
		return tx
	}
}

func filteredAssignee(assigneeID account.ID) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if assigneeID.IsNotEmpty() {
			return tx.Where(
				"EXISTS (SELECT 'e' FROM event_assignment_view ea "+
					"WHERE ea.assignee_id = ? AND ea.event_id = events.id)",
				assigneeID,
			)
		}
		return tx
	}
}

func filteredAttendee(attendeeID account.ID) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if attendeeID.IsNotEmpty() {
			return tx.Where(
				"EXISTS (SELECT 'e' FROM event_booking_view eb "+
					"WHERE eb.attendee_id = ? AND eb.event_id = events.id)",
				attendeeID,
			)
		}
		return tx
	}
}

func orders(orders []string) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		for _, order := range orders {
			tx = tx.Order(order)
		}
		return tx
	}
}

func preloadAssignments(tx *gorm.DB) *gorm.DB {
	return tx.Preload(
		"Assignments", func(db *gorm.DB) *gorm.DB {
			return db.Table("event_assignment_view")
		},
	)
}

func preloadBookings(tx *gorm.DB) *gorm.DB {
	return tx.Preload(
		"Bookings", func(db *gorm.DB) *gorm.DB {
			return db.Table("event_booking_view")
		},
	)
}
