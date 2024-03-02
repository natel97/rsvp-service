package timeselection

import "gorm.io/gorm"

type TimeSelection struct {
	gorm.Model
	ID           string
	InvitationID string
	TimeOptionID string
	Acceptable   *bool
}
