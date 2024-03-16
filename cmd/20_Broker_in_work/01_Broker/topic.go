package main

import (
	"errors"
	"sync"
)

type CodeProducer string
type CodeConsumer string

type Socket string

type topicConfiguration struct {
	secondsTTL int64

	registeredProducers map[CodeProducer]struct{}
	registeredConsumers map[CodeConsumer]Socket
}

type Topic struct {
	topicConfiguration
	mux sync.Mutex

	messages []*Message
}

func NewTopic(secondsTTL int64) *Topic {
	return &Topic{
		topicConfiguration: topicConfiguration{
			secondsTTL:          secondsTTL,
			registeredProducers: make(map[CodeProducer]struct{}),
			registeredConsumers: make(map[CodeConsumer]Socket),
		},

		messages: make([]*Message, 0),
	}
}

// TODO: move to LRU.
func (q *Topic) AddMessage(e *Message) {
	q.mux.Lock()

	q.messages = append(q.messages, e)

	q.mux.Unlock()
}

func (q *Topic) upsertConsumer(c *Consumer) {
	q.registeredConsumers[c.Code] = c
}

func (q *Topic) upsertProducer(p *Producer) {
	q.registeredProducers[p.Code] = p
}

func (q *Topic) getProducer(code string) (*Producer, error) {
	res, exists := q.registeredProducers[code]
	if !exists {
		return nil,
			errors.New("not found")
	}

	return res,
		nil
}

func (q *Topic) getConsumer(code string) (*Consumer, error) {
	res, exists := q.registeredConsumers[code]
	if !exists {
		return nil,
			errors.New("not found")
	}

	return res,
		nil
}
