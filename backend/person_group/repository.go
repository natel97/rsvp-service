package persongroup

import (
	"rsvp/person"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUserGroup(ug *PersonGroup) error
	DeleteUserGroup(groupID string, personID string) error
	GetUsersInGroup(groupID string) ([]InvitedPerson, error)
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *gorm.DB
}

func (repo *repository) CreateUserGroup(ug *PersonGroup) error {
	ug.ID = uuid.New().String()
	err := repo.db.Create(ug).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) DeleteUserGroup(groupID string, personID string) error {
	err := repo.db.Delete(&PersonGroup{}, "group_id = ? AND person_id = ?", groupID, personID).Error

	if err != nil {
		return err
	}

	return nil
}

type InvitedPerson struct {
	InGroup bool
	person.Person
}

func (repo *repository) GetUsersInGroup(groupID string) ([]InvitedPerson, error) {
	people := []InvitedPerson{}

	err := repo.db.Raw(`SELECT p.*,
	CASE WHEN ug.id IS NOT NULL THEN TRUE ELSE FALSE END as in_group
	FROM people p
	LEFT JOIN (
		SELECT * FROM person_groups ug
		WHERE group_id = ?
		AND deleted_at IS NULL
	) ug
	ON p.id = ug.person_id
	`, groupID).Scan(&people).Error

	if err != nil {
		return nil, err
	}
	return people, nil
}
