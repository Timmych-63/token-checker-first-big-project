package model

import "time"

type User struct {
	ID uint `gorm:"primaryKey"`

	Login string `gorm:"size:50;not null;uniqueIndex"`

	PasswordHash string `gorm:"size:255;not null" json:"-"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
