package entities

import "gorm.io/gorm"

type Follower struct {
	gorm.Model
	UserID     uint `gorm:"primaryKey" json:"user_id,omitempty"`
	FollowerID uint `gorm:"primaryKey" json:"follower_id,omitempty"`
}
