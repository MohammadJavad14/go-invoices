package models

import "gorm.io/gorm"

type UserInfo struct {
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	BirthYear uint   `json:"birth_year"`
}
type User struct {
	gorm.Model
	UserInfo
}

type CreateUserInput struct {
	Name      string `json:"name" binding:"required"`
	Gender    string `json:"gender"`
	BirthYear uint   `json:"birth_year"`
}
