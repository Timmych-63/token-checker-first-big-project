package model

import "time"

type Session struct {
	ID uint `gorm:"primaryKey"`

	UserID uint `gorm:"not null;index"`

	TokenHash string `gorm:"size:64;not null;uniqueIndex"`

	ExpiresAt time.Time `gorm:"not null;index"`

	CreatedAt time.Time
}
