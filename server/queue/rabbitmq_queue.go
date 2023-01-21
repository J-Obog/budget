package queue

import (
	"context"
	"encoding/json"

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

func (this *RabbitMqQueue) Push(message Message) error {
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

func (this *RabbitMqQueue) Pull() ([]Message, error) {
	messageChan, err := this.channel.Consume(this.name, "", true, false, false, false, nil)
	messages := make([]Message, len(messageChan))

	if err != nil {
		return nil, err
	}

	for message := range messageChan {
		var msg Message

		err = json.Unmarshal(message.Body, &msg)

		if err != nil {
			return nil, nil
		}

		messages = append(messages, msg)
	}

	return messages, err
}
