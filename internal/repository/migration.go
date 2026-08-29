package repository

import (
	"TOKENCHECKER/internal/model"
	"fmt"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		return fmt.Errorf("не удалось выполнить имграцию базы данный: %w", err)
	}

	return nil
}
