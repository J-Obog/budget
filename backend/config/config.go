package config

import (
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

const (
	LimitMaxAccountNameChars     = 100
	LimitMinAccountNameChars     = 1
	LimitMaxTransactionNoteChars = 200
	LimitMaxCategoryNameChars    = 150
	LimitMinCategoryNameChars    = 1
)

type AppConfig struct {
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresPort     string `env:"POSTGRES_PORT"`
	PostgresDb       string `env:"POSTGRES_DB"`
	RabbitMqUrl      string `env:"RABBIT_MQ_URL"`
	ServerAddress    string `env:"SERVER_ADDRESS"`
	ServerPort       int    `env:"SERVER_PORT"`
}

func Get() *AppConfig {
	environment := os.Getenv("APP_ENV")
	if environment == "dev" {
		if err := godotenv.Load("../.env"); err != nil {
			log.Fatal(err)
		}
	}
	cfg := &AppConfig{}

	if err := env.Parse(cfg); err != nil {
		panic(err)
	}

	return cfg
}
