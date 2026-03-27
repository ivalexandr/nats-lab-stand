package common

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	NATS_URL    string
	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load()

	natsUrl, err := requireField("NATS_URL")
	if err != nil {
		return nil, err
	}
	dbHost, err := requireField("DB_HOST")
	if err != nil {
		return nil, err
	}
	dbPortStr, err := requireField("DB_PORT")
	if err != nil {
		return nil, err
	}
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return nil, err
	}
	dbUser, err := requireField("DB_USER")
	if err != nil {
		return nil, err
	}
	dbPassword, err := requireField("DB_PASSWORD")
	if err != nil {
		return nil, err
	}
	dbName, err := requireField("DB_NAME")
	if err != nil {
		return nil, err
	}

	return &Config{
		NATS_URL:    natsUrl,
		DB_HOST:     dbHost,
		DB_PORT:     dbPort,
		DB_USER:     dbUser,
		DB_PASSWORD: dbPassword,
		DB_NAME:     dbName,
	}, nil
}

func requireField(name string) (string, error) {
	value := os.Getenv(name)
	if value == "" {
		return "", fmt.Errorf("env %s is required", name)
	}

	return value, nil
}
