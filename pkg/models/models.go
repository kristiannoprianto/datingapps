package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Premium  bool   `gorm:"default:false"`
}

type Profile struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	Name        string `gorm:"not null"`
	Gender      string
	DateOfBirth time.Time
	Interest    string `json:"description"`
}

type Swipe struct {
	gorm.Model
	UserID    uint
	ProfileID uint
	Action    string // "like" or "pass"
}
