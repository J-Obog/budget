package queue

import (
	"log"

	"github.com/J-Obog/paidoff/config"
)

const (
	queueImpl = "rabbit"
)

func CreateConfig(app *config.AppConfig) Queue {
	switch queueImpl {
	case "rabbit":
		return getRabbitMqQueue(app.RabbitMqUrl)
	default:
		log.Fatal("Not a supported impl for queue")
	}

	return nil
}

func getRabbitMqQueue(url string) *RabbitMqQueue {
	return NewRabbitMqQueue(url)
}
