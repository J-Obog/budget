package config

import (
	"errors"
	"os"

	"github.com/J-Obog/paidoff/store"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	appStoreImpl = "postgres"
)

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

func MakeStore(cfg *AppConfig) (store.Store, error) {
	switch appStoreImpl {
	case "postgres":
		pgDb, err := gorm.Open(postgres.Open(cfg.PostgresUrl), &gorm.Config{})

		if err != nil {
			return nil, err
		}

		return store.NewPostgresStore(pgDb), nil

	default:
		return nil, errors.New("cannot find impl for store")
	}
}
