package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("не удалось загрузить файл .env: %w", err)
	}

	cfg := &Config{
		AppPort:    envOrDefault("APP_PORT", "8080"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
	}

	requiredVariables := []struct {
		name  string
		value string
	}{
		{name: "DB_HOST", value: cfg.DBHost},
		{name: "DB_PORT", value: cfg.DBPort},
		{name: "DB_USER", value: cfg.DBUser},
		{name: "DB_PASSWORD", value: cfg.DBPassword},
		{name: "DB_NAME", value: cfg.DBName},
	}
	var missingVariables []string

	for _, variable := range requiredVariables {
		if strings.TrimSpace(variable.value) == "" {
			missingVariables = append(missingVariables, variable.name)
		}
	}

	if len(missingVariables) > 0 {
		return nil, fmt.Errorf(
			"не заданы обязательные переменные окружения %s",
			strings.Join(missingVariables, ", "),
		)
	}

	return cfg, nil
}

func envOrDefault(key string, defaultValue string) string {
	value := strings.TrimSpace(os.Getenv(key))

	if value == "" {
		return defaultValue
	}

	return value
}
