package person

import "gorm.io/gorm"

type Person struct {
	gorm.Model

	ID    string
	First string
	Last  string
}
