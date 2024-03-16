package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMessageCreation(t *testing.T) {
	producer := Producer{
		Code: "P1",
	}

	topic := NewTopic(100)

	message := NewMessage(
		&ParamsNewMessage{
			CodeProducer: producer.Code,
		},
		[]byte("payload"),
	)

	topic.AddMessage(message)
	require.Len(t, topic.messages, 1)
}

func TestGetEvent(t *testing.T) {
	producer := Producer{
		Code: "P1",
	}

	topic := NewTopic(100)

	topic.AddMessage(
		NewMessage(
			&ParamsNewMessage{
				CodeProducer: producer.Code,
			},
			[]byte("payload1"),
		),
	)
	topic.AddMessage(
		NewMessage(
			&ParamsNewMessage{
				CodeProducer: producer.Code,
			},
			[]byte("payload2"),
		),
	)

	// TODO: push to consumer.
	c1 := Consumer{
		Code:   "C1",
		Socket: "127.0.0.2:8081",
	}
	require.NotZero(t, c1)
}
