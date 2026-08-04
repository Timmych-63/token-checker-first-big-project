package repository

import (
	"fmt"
	"net"
	"net/url"

	"TOKENCHECKER/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := buildDatabaseURL(cfg)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf(
			"не удалось открыть подключение к PostgreSQL: %w",
			err,
		)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf(
			"не удалось получить низкоуровневое подключение к базе: %w",
			err,
		)
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf(
			"PostgreSQL не отвечает: %w",
			err,
		)
	}

	return db, nil

}

func buildDatabaseURL(cfg *config.Config) string {
	databaseURL := &url.URL{
		Scheme: "postgres",

		User: url.UserPassword(
			cfg.DBUser,
			cfg.DBPassword,
		),

		Host: net.JoinHostPort(
			cfg.DBHost,
			cfg.DBPort,
		),

		Path: cfg.DBName,
	}

	query := databaseURL.Query()
	query.Set("sslmode", cfg.DBSSLMode)

	databaseURL.RawQuery = query.Encode()

	return databaseURL.String()
}
