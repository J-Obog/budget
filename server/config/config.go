package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	LimitMaxAccountNameChars = 100
	LimitMinAccountNameChars = 1

	LimitMaxTransactionNoteChars = 200

	LimitMaxCategoryNameChars = 150
	LimitMinCategoryNameChars = 1
)

type AppConfig struct {
	PostgresUrl string `json:"postgresUrl"`
	RabbitMqUrl string `json:"rabbitMqUrl"`
}

func Get() *AppConfig {
	env := os.Getenv("APP_ENV")
	if env == "dev" {
		if err := godotenv.Load("./.env"); err != nil {
			log.Fatal(err)
		}
	}

	return &AppConfig{
		PostgresUrl: os.Getenv("POSTGRES_URL"),
		RabbitMqUrl: os.Getenv("RABBIT_MQ_URL"),
	}
}
