package group

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

type Repository interface {
	Create(e Group) (*Group, error)
	Get(id string) (*Group, error)
	GetAll() ([]Group, error)
	Delete(id string) (*Group, error)
	Update(id string, details Group) (*Group, error)
}

func (repo *repository) Create(e Group) (*Group, error) {
	id := uuid.New()
	e.ID = id.String()

	err := repo.db.Create(&e).Error

	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (repo *repository) Get(id string) (*Group, error) {
	group := Group{}

	err := repo.db.First(&group, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (repo *repository) GetAll() ([]Group, error) {
	group := []Group{}

	err := repo.db.Find(&group).Error
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (repo *repository) Delete(id string) (*Group, error) {
	group := Group{}

	err := repo.db.Delete(&group, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &group, nil
}

type InvitedGroup struct {
	InvitationID string
	Group
}

func (repo *repository) Update(id string, details Group) (*Group, error) {
	details.ID = ""

	match := Group{}

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
