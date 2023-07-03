package queue

import (
	"context"
	"encoding/json"

	"github.com/J-Obog/paidoff/data"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMqQueue struct {
	channel *amqp.Channel
	dtags   map[string]uint64
	name    string
}

func NewRabbitMqQueue(channel *amqp.Channel, name string) *RabbitMqQueue {
	return &RabbitMqQueue{
		channel: channel,
		name:    name,
		dtags:   make(map[string]uint64),
	}
}

func (mq *RabbitMqQueue) Push(message data.Message) error {
	ctx := context.Background()
	payload, err := json.Marshal(message)

	if err != nil {
		return err
	}

	msg := amqp.Publishing{
		MessageId: message.Id,
		Body:      payload,
	}

	return mq.channel.PublishWithContext(ctx, "", mq.name, true, false, msg)
}

func (mq *RabbitMqQueue) Pop() (*data.Message, error) {
	d, ok, err := mq.channel.Get(mq.name, false)

	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, nil
	}

	mq.dtags[d.MessageId] = d.DeliveryTag

	var message = &data.Message{}
	err = json.Unmarshal(d.Body, message)

	if err != nil {
		return nil, err
	}

	return message, err
}
