package invitation

import (
	. "rsvp/invitation/types"

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

func (repo *repository) Create(e Invitation) (*Invitation, error) {
	id := uuid.New()

	err := repo.db.Where(e).Attrs(Invitation{ID: id.String()}).FirstOrCreate(&e).Error

	if err != nil {
		return nil, err
	}

	return &e, nil
}

type UserID string

func (repo *repository) InviteGroup(e InviteGroup) error {
	personIDs := []UserID{}

	err := repo.db.Table("person_groups").Select("person_id").Find(&personIDs, "group_id = ? AND deleted_at IS NULL", e.GroupID).Error

	if err != nil {
		return err
	}

	for _, personID := range personIDs {
		_, err := repo.Create(Invitation{
			PersonID: string(personID),
			EventID:  e.EventID,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *repository) Get(id string) (*Invitation, error) {
	invitation := Invitation{}

	err := repo.db.First(&invitation, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &invitation, nil
}

func (repo *repository) GetAll() ([]Invitation, error) {
	invitation := []Invitation{}

	err := repo.db.Find(&invitation).Error
	if err != nil {
		return nil, err
	}

	return invitation, nil
}

func (repo *repository) GetByEvent(eventID string) ([]Invitation, error) {
	invitation := []Invitation{}

	err := repo.db.Find(&invitation, "event_id = ?", eventID).Error
	if err != nil {
		return nil, err
	}

	return invitation, nil
}

func (repo *repository) Delete(id string) error {
	invitation := Invitation{}

	err := repo.db.Delete(&invitation, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) Update(id string, details Invitation) (*Invitation, error) {
	details.ID = ""

	match := Invitation{}

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
