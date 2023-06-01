package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name,omitempty"`
	Nick     string `json:"nick,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
