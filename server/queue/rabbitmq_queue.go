package queue

import (
	"context"
	"encoding/json"

	"github.com/J-Obog/paidoff/data"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	exchange string = "direct"
)

type RabbitMqQueue struct {
	channel *amqp.Channel
	name    string
}

func NewRabbitMqQueue(channel *amqp.Channel, name string) *RabbitMqQueue {
	return &RabbitMqQueue{
		channel: channel,
		name:    name,
	}
}

func (this *RabbitMqQueue) Push(message data.Message) error {
	ctx := context.Background()

	payload, err := json.Marshal(message)

	if err != nil {
		return err
	}

	msg := amqp.Publishing{
		Body: payload,
	}

	return this.channel.PublishWithContext(ctx, exchange, this.name, true, false, msg)
}

func (this *RabbitMqQueue) Pull() ([]data.Message, error) {
	messageChan, err := this.channel.Consume(this.name, "", true, false, false, false, nil)
	messages := make([]data.Message, len(messageChan))

	if err != nil {
		return nil, err
	}

	for message := range messageChan {
		var msg data.Message

		err = json.Unmarshal(message.Body, &msg)

		if err != nil {
			return nil, nil
		}

		messages = append(messages, msg)
	}

	return messages, err
}
