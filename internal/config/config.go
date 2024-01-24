package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBConfig
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return nil, errors.New("DB_HOST is required")
	}

	dbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		return nil, errors.New("DB_USER is required")
	}

	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return nil, errors.New("DB_PASSWORD is required")
	}

	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return nil, errors.New("DB_NAME is required")
	}

	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return nil, errors.New("DB_PORT is required")
	}

	return &Config{
		DBConfig: DBConfig{
			Host:     dbHost,
			User:     dbUser,
			Password: dbPassword,
			DBName:   dbName,
			Port:     dbPort,
		},
	}, nil
}
