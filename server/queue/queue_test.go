package queue

import (
	"testing"

	"github.com/J-Obog/paidoff/config"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	g, _ := config.MakeConfig(config.EnvType_LOCAL)
	q := MakeQueue(g)

	t.Run("it pushes and pops message", func(t *testing.T) {
		msg := testMessage()

		err := q.Push(msg)
		assert.NoError(t, err)

		m, err := q.Pop()
		assert.NoError(t, err)
		assert.Equal(t, msg, *m)
	})

	t.Run("it acks a message", func(t *testing.T) {
		msg := testMessage()

		err := q.Push(msg)
		assert.NoError(t, err)

		m, err := q.Pop()
		assert.NoError(t, err)

		err = q.Ack(m.Id)
		assert.NoError(t, err)

		msgs, err := q.Pop()
		assert.NoError(t, err)
		assert.Nil(t, msgs)
	})
}
