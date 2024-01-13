package persongroup

import "gorm.io/gorm"

type PersonGroup struct {
	gorm.Model
	ID       string
	PersonID string
	GroupID  string
}
