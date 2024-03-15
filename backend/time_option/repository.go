package timeoption

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

//go:generate mockgen -source=repository.go -destination=repository_mock.go -package=timeoption
type Repository interface {
	CreateTimeOption(to *TimeOption) error
	DeleteTimeOption(id string) error
	GetTimeOption(id string) (*TimeOption, error)
}

func (repo *repository) CreateTimeOption(to *TimeOption) error {
	to.ID = uuid.New().String()
	err := repo.db.Create(to).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) DeleteTimeOption(id string) error {
	err := repo.db.Delete(&TimeOption{}, "id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) GetTimeOption(id string) (*TimeOption, error) {
	to := TimeOption{}
	err := repo.db.Find(&to, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &to, nil
}
