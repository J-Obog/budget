package queue

import (
	"testing"

	"github.com/J-Obog/paidoff/config"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	g, _ := config.MakeConfig(config.EnvType_LOCAL)

	q := MakeQueue(g)

	err := q.Flush(testQueueName)
	assert.NoError(t, err)

	t.Run("it pushes and pops message and acks message", func(t *testing.T) {
		msg := testMessage()

		err := q.Push(msg, testQueueName)
		assert.NoError(t, err)

		m, err := q.Pop(testQueueName)
		assert.NoError(t, err)
		assert.Equal(t, msg, *m)

		err = q.Ack(m.Id)
		assert.NoError(t, err)

		msgs, err := q.Pop(testQueueName)
		assert.NoError(t, err)
		assert.Nil(t, msgs)
	})
}
