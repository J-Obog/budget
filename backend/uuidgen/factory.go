package uuid

import (
	"log"

	"github.com/J-Obog/paidoff/config"
)

const (
	uuidImpl = "ksuid"
)

func NewUuidProvider(app *config.AppConfig) UuidProvider {
	switch uuidImpl {
	case "ksuid":
		return NewKsuidProvider()
	default:
		log.Fatal("Not a supported impl for clock")
	}

	return nil
}
