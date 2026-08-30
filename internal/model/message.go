package model

import "time"

type Message struct {
	ID uint `gorm:"primaryKey"`

	UserID uint `gorm:"not null;uniqueIndex"`

	Text string `gorm:"type:text;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SaveMessageRequest struct {
	Text string `json:"text"`
}

type SaveMessageResponse struct {
	Message string `json:"message"`
}

type GetMessageResponse struct {
	Text string `json:"text"`
}
