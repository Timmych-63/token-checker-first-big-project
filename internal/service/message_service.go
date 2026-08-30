package service

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"TOKENCHECKER/internal/repository"
)

var (
	ErrMessageRequired = errors.New(
		"сообщение не может быть пустым",
	)

	ErrMessageTooLong = errors.New(
		"сообщение должно содержать не более 5000 символов",
	)
)

type MessageService struct {
	messageRepository *repository.MessageRepository
}

func NewMessageService(
	messageRepository *repository.MessageRepository,
) *MessageService {
	return &MessageService{
		messageRepository: messageRepository,
	}
}

func (s *MessageService) SaveMessage(
	userID uint,
	text string,
) error {
	if strings.TrimSpace(text) == "" {
		return ErrMessageRequired
	}

	if utf8.RuneCountInString(text) > 5000 {
		return ErrMessageTooLong
	}

	err := s.messageRepository.SaveForUser(
		userID,
		text,
	)
	if err != nil {
		return fmt.Errorf(
			"не удалось сохранить сообщение пользователя: %w",
			err,
		)
	}

	return nil
}

func (s *MessageService) GetMessage(
	userID uint,
) (string, error) {
	message, err := s.messageRepository.FindByUserID(
		userID,
	)
	if err != nil {
		return "", fmt.Errorf(
			"не удалось загрузить сообщение пользователя: %w",
			err,
		)
	}

	if message == nil {
		return "", nil
	}

	return message.Text, nil
}
