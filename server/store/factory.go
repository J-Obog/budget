package store

import (
	"errors"

	"github.com/J-Obog/paidoff/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	storeImpl = "postgres"
)

func MakeStore(cfg *config.AppConfig) (Store, error) {
	switch storeImpl {
	case "postgres":
		pgDb, err := gorm.Open(postgres.Open(cfg.PostgresUrl), &gorm.Config{
			AllowGlobalUpdate: true,
			NowFunc:           nil,
			Logger:            logger.Default.LogMode(logger.Silent),
		})

		if err != nil {
			return nil, err
		}

		return NewPostgresStore(pgDb), nil

	default:
		return nil, errors.New("cannot find impl for store")
	}
}
