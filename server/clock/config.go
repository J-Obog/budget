package clock

import (
	"log"

	"github.com/J-Obog/paidoff/config"
)

const (
	clockImpl = "system"
)

func CreateConfig(app *config.AppConfig) Clock {
	switch clockImpl {
	case "system":
		return NewSystemClock()
	default:
		log.Fatal("Not a supported impl for clock")
	}

	return nil
}
