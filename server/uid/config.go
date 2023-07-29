package uid

import (
	"log"

	"github.com/J-Obog/paidoff/config"
)

const (
	uuidImpl = "ksuid"
)

func CreateConfig(app *config.AppConfig) UIDProvider {
	switch uuidImpl {
	case "ksuid":
		return NewKSUIDProvider()
	default:
		log.Fatal("Not a supported impl for clock")
	}

	return nil
}
