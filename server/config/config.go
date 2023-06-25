package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	PostgresUrl string `json:"postgresUrl"`
}

func MakeConfig(env string) (*AppConfig, error) {

	if env == "local" {
		if err := godotenv.Load(); err != nil {
			return nil, err
		}
	}

	return &AppConfig{
		PostgresUrl: os.Getenv("POSTGRES_URL"),
	}, nil
}
