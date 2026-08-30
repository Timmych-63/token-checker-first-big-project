package repository

import (
	"fmt"

	"TOKENCHECKER/internal/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.User{},
		&model.Session{},
		&model.Message{},
	)
	if err != nil {
		return fmt.Errorf(
			"не удалось выполнить миграцию базы данных: %w",
			err,
		)
	}

	return nil
}
