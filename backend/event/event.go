package event

import (
	"time"

	"gorm.io/gorm"
)

type EventState string

const (
	PLANNING            = "PLANNING"
	UPCOMING            = "UPCOMING"
	CANCELED            = "CANCELED"
	STALE    EventState = "DONE"
)

var states = [...]EventState{PLANNING, UPCOMING, CANCELED, STALE}

func GetStateID(state EventState) uint {
	for id, val := range states {
		if val == state {
			return uint(id)
		}
	}
	return 0
}

type Event struct {
	gorm.Model

	ID           string
	Title        string
	Date         *time.Time
	Street       string
	City         string
	Description  string
	InternalNote string
	State        uint
}

type TimeOption struct {
	ID         string     `json:"id"`
	Time       *time.Time `json:"time"`
	Upvote     uint       `json:"upvote"`
	Downvote   uint       `json:"downvote"`
	IsUpvote   bool       `json:"isUpvote"`
	IsDownvote bool       `json:"isDownvote"`
}

func (e *Event) GetEventState() string {
	return string(states[e.State])
}

func (e *Event) SetEventState(state EventState) {
	id := GetStateID(state)
	e.State = id
}
