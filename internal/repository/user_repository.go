package repository

import (
	"errors"
	"fmt"

	"TOKENCHECKER/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *model.User) error {
	result := r.db.Create(user)

	if result.Error != nil {
		return fmt.Errorf(
			"не удалось сохранить пользователя: %w",
			result.Error,
		)
	}

	return nil
}

func (r *UserRepository) FindByLogin(login string) (*model.User, error) {
	var user model.User

	result := r.db.
		Where("login = ?", login).
		First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, fmt.Errorf(
			"не удалось найти пользователя по логину: %w",
			result.Error,
		)
	}

	return &user, nil
}
