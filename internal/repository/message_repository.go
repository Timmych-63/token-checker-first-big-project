package repository

import (
	"errors"
	"fmt"

	"TOKENCHECKER/internal/model"

	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(
	db *gorm.DB,
) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

func (r *MessageRepository) SaveForUser(
	userID uint,
	text string,
) error {
	var message model.Message

	result := r.db.
		Where("user_id = ?", userID).
		First(&message)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		message = model.Message{
			UserID: userID,
			Text:   text,
		}

		result = r.db.Create(&message)

		if result.Error != nil {
			return fmt.Errorf(
				"не удалось создать сообщение: %w",
				result.Error,
			)
		}

		return nil
	}

	if result.Error != nil {
		return fmt.Errorf(
			"не удалось найти сообщение перед сохранением: %w",
			result.Error,
		)
	}

	message.Text = text

	result = r.db.Save(&message)

	if result.Error != nil {
		return fmt.Errorf(
			"не удалось обновить сообщение: %w",
			result.Error,
		)
	}

	return nil
}

func (r *MessageRepository) FindByUserID(
	userID uint,
) (*model.Message, error) {
	var message model.Message

	result := r.db.
		Where("user_id = ?", userID).
		First(&message)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, fmt.Errorf(
			"не удалось получить сообщение пользователя: %w",
			result.Error,
		)
	}

	return &message, nil
}
