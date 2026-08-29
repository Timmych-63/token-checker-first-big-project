package repository

import (
	"errors"
	"fmt"

	"TOKENCHECKER/internal/model"

	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

func (r *SessionRepository) Create(
	session *model.Session,
) error {
	result := r.db.Create(session)

	if result.Error != nil {
		return fmt.Errorf(
			"не удалось сохранить сессию: %w",
			result.Error,
		)
	}

	return nil
}

func (r *SessionRepository) FindByTokenHash(
	tokenHash string,
) (*model.Session, error) {
	var session model.Session

	result := r.db.
		Where("token_hash = ?", tokenHash).
		First(&session)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, fmt.Errorf(
			"не удалось найти сессию: %w",
			result.Error,
		)
	}

	return &session, nil
}

func (r *SessionRepository) DeleteByTokenHash(
	tokenHash string,
) error {
	result := r.db.
		Where("token_hash = ?", tokenHash).
		Delete(&model.Session{})

	if result.Error != nil {
		return fmt.Errorf(
			"не удалось удалить сессию: %w",
			result.Error,
		)
	}

	return nil
}
