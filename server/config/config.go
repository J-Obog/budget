package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvType uint

const (
	EnvType_LOCAL EnvType = 0
)

const (
	LimitMaxAccountNameChars     = 100
	LimitMinAccountNameChars     = 1
	LimitMaxTransactionNoteChars = 200
	LimitMaxCategoryNameChars    = 150
)

type AppConfig struct {
	PostgresUrl string `json:"postgresUrl"`
	RabbitMqUrl string `json:"rabbitMqUrl"`
}

func MakeConfig(envType EnvType) (*AppConfig, error) {
	if envType == EnvType_LOCAL {
		if err := godotenv.Load("../.env.local"); err != nil {
			return nil, err
		}
	}

	return &AppConfig{
		PostgresUrl: os.Getenv("POSTGRES_URL"),
		RabbitMqUrl: os.Getenv("RABBIT_MQ_URL"),
	}, nil
}
