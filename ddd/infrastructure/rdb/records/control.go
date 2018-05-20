package records

type Control struct {
	EventID uint `gorm:"primary_key"`
	Version uint `gorm:"not null"`
}

func (*Control) TableName() string {
	return "event_controls"
}
