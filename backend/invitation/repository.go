package invitation

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

func (repo *repository) Create(e Invitation) (*Invitation, error) {
	id := uuid.New()
	e.ID = id.String()

	err := repo.db.Create(&e).Error

	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (repo *repository) Get(id string) (*Invitation, error) {
	invitation := Invitation{}

	err := repo.db.Find(&invitation, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &invitation, nil
}

func (repo *repository) Delete(id string) (*Invitation, error) {
	invitation := Invitation{}

	err := repo.db.Delete(&invitation, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &invitation, nil
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
