package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	BirthYear uint   `json:"birth_year"`
}
