package rsvp

import "gorm.io/gorm"

// TODO: Break out Going/BringingFriend from struct
type RSVP struct {
	gorm.Model

	ID             string
	InvitationID   string
	Going          string
	BringingFriend string
}
