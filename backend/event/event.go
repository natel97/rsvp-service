package event

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model

	ID     string
	Title  string
	Date   *time.Time
	Street string
	City   string
}
