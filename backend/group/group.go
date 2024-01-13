package group

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	ID          string
	Name        string
	Description string
}
