package record

type Control struct {
	EventID string `gorm:"primary_key"`
	Version uint   `gorm:"not null"`
}

func (*Control) TableName() string {
	return "event_controls"
}
