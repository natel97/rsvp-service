package invitation

import (
	"time"
)

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
	Description  string     `json:"description"`
	Subscribed   bool       `json:"subscribed"`
}
