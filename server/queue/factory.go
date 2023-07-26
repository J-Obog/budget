package queue

import (
	"log"

	"github.com/J-Obog/paidoff/config"
)

const (
	queueImpl = "rabbit"
)

func GetConfiguredQueue(cfg *config.AppConfig) Queue {
	switch queueImpl {
	case "rabbit":
		return NewRabbitMqQueue(cfg.RabbitMqUrl)
	default:
		log.Fatal("Not a supported impl for queue")
	}

	return nil
}
