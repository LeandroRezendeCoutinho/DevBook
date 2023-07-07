package entities

import "gorm.io/gorm"

type Follower struct {
	gorm.Model
	UserID      uint `gorm:"primaryKey" json:"user_id,omitempty"`
	FollowingID uint `gorm:"primaryKey" json:"following_id,omitempty"`
}
