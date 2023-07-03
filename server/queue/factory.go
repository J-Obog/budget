package queue

import (
	"log"

	"github.com/J-Obog/paidoff/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	qImpl = "rabbit"
)

func MakeQueue(cfg *config.AppConfig) Queue {
	switch qImpl {
	case "rabbit":
		conn, err := amqp.Dial(cfg.RabbitMqUrl)
		if err != nil {
			log.Fatal(err)
		}

		ch, err := conn.Channel()
		ch.Qos(1, 0, false)
		ch.Confirm(false)

		if err != nil {
			log.Fatal(err)
		}

		return NewRabbitMqQueue(ch, "foob")

	default:
		log.Fatal("Not a supported impl for queue")
	}

	return nil
}
