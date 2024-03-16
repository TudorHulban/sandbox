package main

import (
	"time"

	"github.com/TudorHulban/epochid"
)

type MessageHeaders struct {
	CodeProducer string
	Consumers    map[CodeConsumer]Socket

	ID                  uint64
	CreatedAt           int64
	RequestedTTLSeconds int64 // if value, enforce TTL
}

type Message struct {
	MessageHeaders

	Body []byte
}

type ParamsNewMessage struct {
	CodeProducer string
	Consumers    map[CodeConsumer]Socket

	RequestedTTLSeconds int64
}

func NewMessage(params *ParamsNewMessage, body []byte) *Message {
	return &Message{
		MessageHeaders: MessageHeaders{
			ID:                  epochid.NewIDIncremental10KWCoCorrection(),
			CreatedAt:           time.Now().UnixNano(),
			RequestedTTLSeconds: params.RequestedTTLSeconds,

			Consumers:    params.Consumers,
			CodeProducer: params.CodeProducer,
		},

		Body: body,
	}
}

// once consumer accepts push, it is deleted from consumers.
func (m *Message) PushMessage() error {
	return nil
}
