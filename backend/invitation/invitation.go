package invitation

import (
	"time"

	"gorm.io/gorm"
)

type Invitation struct {
	gorm.Model

	ID      string
	UserID  string
	EventID string
}

type Attendance struct {
	Yes   uint `json:"yes"`
	No    uint `json:"no"`
	Maybe uint `json:"maybe"`
}

type GetInvitationResponse struct {
	Title        string     `json:"title"`
	Date         *time.Time `json:"date"`
	Street       string     `json:"street"`
	City         string     `json:"city"`
	Attendance   Attendance `json:"attendance"`
	MyAttendance string     `json:"myAttendance"`
	MyFriend     string     `json:"myFriend"`
}
