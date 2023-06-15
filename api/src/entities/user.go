package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name,omitempty" validate:"required"`
	Nick     string `json:"nick,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=6,max=12"`
}
