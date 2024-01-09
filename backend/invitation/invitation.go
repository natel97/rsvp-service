package invitation

import "gorm.io/gorm"

type Invitation struct {
	gorm.Model

	ID      string
	UserID  string
	EventID string
}
