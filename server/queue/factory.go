package queue

import (
	"log"

	"github.com/J-Obog/paidoff/config"
)

const (
	queueImpl = "rabbit"
)

func NewQueue(app *config.AppConfig) Queue {
	switch queueImpl {
	case "rabbit":
		return NewRabbitMqQueue(app.RabbitMqUrl)
	default:
		log.Fatal("Not a supported impl for queue")
	}

	return nil
}
