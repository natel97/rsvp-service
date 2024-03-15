package types

import "gorm.io/gorm"

type Invitation struct {
	gorm.Model

	ID       string
	PersonID string
	EventID  string
}

type InviteGroup struct {
	GroupID string
	EventID string
}

//go:generate mockgen -source=types.go -destination=types_mock.go -package=types
type Repository interface {
	Create(e Invitation) (*Invitation, error)
	Get(id string) (*Invitation, error)
	GetAll() ([]Invitation, error)
	GetByEvent(eventID string) ([]Invitation, error)
	Delete(id string) error
	Update(id string, details Invitation) (*Invitation, error)
	InviteGroup(e InviteGroup) error
}
