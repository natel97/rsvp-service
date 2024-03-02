package timeoption

import (
	"time"

	"gorm.io/gorm"
)

type TimeOption struct {
	gorm.Model
	ID      string
	EventID string
	Time    *time.Time
}
