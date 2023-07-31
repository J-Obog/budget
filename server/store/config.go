package store

import (
	"log"

	"github.com/J-Obog/paidoff/config"
)

const (
	storeImpl = "postgres"
)

func CreateConfig(app *config.AppConfig) Store {
	switch storeImpl {
	case "postgres":
		return NewPostgresStore(app.PostgresUrl)
	default:
		log.Fatal("Not a supported impl for store")
	}

	return nil
}
