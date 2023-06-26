package store

import (
	"errors"

	"github.com/J-Obog/paidoff/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	storeImpl = "postgres"
)

func MakeStore(cfg *config.AppConfig) (Store, error) {
	switch storeImpl {
	case "postgres":
		pgDb, err := gorm.Open(postgres.Open(cfg.PostgresUrl), &gorm.Config{
			AllowGlobalUpdate: true,
		})

		if err != nil {
			return nil, err
		}

		return NewPostgresStore(pgDb), nil

	default:
		return nil, errors.New("cannot find impl for store")
	}
}
