package rsvp

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) Create(e RSVP) (*RSVP, error) {
	id := uuid.New()
	e.ID = id.String()

	err := repo.db.Create(&e).Error

	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (repo *repository) GetEventRSVPs(eventID string) ([]RSVP, error) {
	rsvp := []RSVP{}

	err := repo.db.Raw(`
	SELECT r.* FROM (
		SELECT max(created_at) created_at, invitation_id
	    FROM rsvps
		WHERE event_id = ?
		GROUP BY invitation_id
	) AS newest
	INNER JOIN rsvps r
	ON newest.created_at = r.created_at
		AND newest.invitation_id = r.invitation_id`, eventID).
		Scan(&rsvp).Error
	if err != nil {
		return nil, err
	}

	return rsvp, nil
}

func (repo *repository) Get(id string) (*RSVP, error) {
	rsvp := RSVP{}

	err := repo.db.Find(&rsvp, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &rsvp, nil
}

func (repo *repository) GetLatestRSVPByInvitation(invitationID string) (*RSVP, error) {
	rsvp := RSVP{}
	err := repo.db.Order("created_at DESC").First(&rsvp, "invitation_id = ?", invitationID).Error
	if err != nil {
		return nil, err
	}

	return &rsvp, nil
}

func (repo *repository) Delete(id string) (*RSVP, error) {
	rsvp := RSVP{}

	err := repo.db.Delete(&rsvp, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &rsvp, nil
}

func (repo *repository) Update(id string, details RSVP) (*RSVP, error) {
	details.ID = ""

	match := RSVP{}

	err := repo.db.First(&match, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	err = repo.db.Where("id = ?", id).Updates(&details).Error

	if err != nil {
		return nil, err
	}

	return &details, nil
}
