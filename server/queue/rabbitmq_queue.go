package queue

import (
	"context"
	"encoding/json"
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMqQueue struct {
	channel *amqp.Channel
	dtags   map[string]uint64
}

func NewRabbitMqQueue(channel *amqp.Channel) *RabbitMqQueue {
	return &RabbitMqQueue{
		channel: channel,
		dtags:   make(map[string]uint64),
	}
}

func (mq *RabbitMqQueue) Push(message Message, queueName string) error {
	ctx := context.Background()
	payload, err := json.Marshal(message)

	dtag := mq.channel.GetNextPublishSeqNo()

	mq.dtags[message.Id] = dtag

	if err != nil {
		return err
	}

	msg := amqp.Publishing{
		Body: payload,
	}

	return mq.channel.PublishWithContext(ctx, "", queueName, true, false, msg)
}

func (mq *RabbitMqQueue) Pop(queueName string) (*Message, error) {

	d, ok, err := mq.channel.Get(queueName, false)

	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, nil
	}

	message := &Message{}
	err = json.Unmarshal(d.Body, message)

	if err != nil {
		return nil, err
	}

	return message, err
}

func (mq *RabbitMqQueue) Ack(messageId string) error {
	tag, ok := mq.dtags[messageId]
	delete(mq.dtags, messageId)

	if !ok {
		//update error message
		return errors.New("some errors")
	}

	return mq.channel.Ack(tag, false)
}
