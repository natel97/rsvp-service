package types

import "gorm.io/gorm"

type Invitation struct {
	gorm.Model

	ID       string
	PersonID string
	EventID  string
}

type Repository interface {
	Create(e Invitation) (*Invitation, error)
	Get(id string) (*Invitation, error)
	GetAll() ([]Invitation, error)
	GetByEvent(eventID string) ([]Invitation, error)
	Delete(id string) error
	Update(id string, details Invitation) (*Invitation, error)
}
