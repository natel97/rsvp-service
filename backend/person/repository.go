package person

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
	Create(e Person) (*Person, error)
	Get(id string) (*Person, error)
	GetAll() ([]Person, error)
	Delete(id string) (*Person, error)
	Update(id string, details Person) (*Person, error)
}

func (repo *repository) Create(e Person) (*Person, error) {
	id := uuid.New()
	e.ID = id.String()

	err := repo.db.Create(&e).Error

	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (repo *repository) Get(id string) (*Person, error) {
	person := Person{}

	err := repo.db.Find(&person, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &person, nil
}

func (repo *repository) GetAll() ([]Person, error) {
	person := []Person{}

	err := repo.db.Find(&person).Error
	if err != nil {
		return nil, err
	}

	return person, nil
}

func (repo *repository) Delete(id string) (*Person, error) {
	person := Person{}

	err := repo.db.Delete(&person, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &person, nil
}

func (repo *repository) Update(id string, details Person) (*Person, error) {
	details.ID = ""

	match := Person{}

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
