package timeselection

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
	UpdateSelection(to *TimeSelection) error
	DeleteSelection(id string) error
}

func (repo *repository) UpdateSelection(ts *TimeSelection) error {
	id := uuid.New().String()

	value := TimeSelection{
		InvitationID: ts.InvitationID,
		TimeOptionID: ts.TimeOptionID,
		Acceptable:   ts.Acceptable,
	}
	err := repo.db.
		Where("invitation_id = ? AND time_option_id = ?", ts.InvitationID, ts.TimeOptionID).
		Attrs(TimeSelection{ID: id, InvitationID: ts.InvitationID, TimeOptionID: ts.TimeOptionID}).
		Assign(map[string]interface{}{"acceptable": ts.Acceptable}).
		FirstOrCreate(&value).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) DeleteSelection(id string) error {
	err := repo.db.Delete(&TimeSelection{}, "id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}
