package event

import (
	"time"
)

type Event struct {
	ID     string
	Title  string
	Date   *time.Time
	Street string
	City   string
}
