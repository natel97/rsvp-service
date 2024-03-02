package rsvp

import "gorm.io/gorm"

type UpdateRSVP struct {
	Going          string `json:"going"`
	BringingFriend string `json:"bringingAFriend"`
}

// TODO: Break out Going/BringingFriend from struct
type RSVP struct {
	gorm.Model

	ID             string
	InvitationID   string
	EventID        string
	Going          string
	BringingFriend string
}

type Repository interface {
	Create(e RSVP) (*RSVP, error)
	GetEventRSVPs(eventID string) ([]RSVP, error)
	GetLatestRSVPByInvitation(invitationID string) (*RSVP, error)
	Get(id string) (*RSVP, error)
	Delete(id string) (*RSVP, error)
	Update(id string, details RSVP) (*RSVP, error)
}
