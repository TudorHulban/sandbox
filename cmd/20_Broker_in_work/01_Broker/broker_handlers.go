package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
)

func (b *Broker) handlerRegisterProducer(c *fiber.Ctx) error {
	req, errDecode := decodeRequest[Producer](c.Request())
	if errDecode != nil {

		return sendResponse(c, errDecode.Error())
	}

	var generatedProducer Producer

	mapstructure.Decode(req, &generatedProducer)

	b.Queues[0].upsertProducer(&generatedProducer)

	return sendResponse(c, "ok")
}

func (b *Broker) handlerRegisterConsumer(c *fiber.Ctx) error {
	req, errDecode := decodeRequest[Consumer](c.Request())
	if errDecode != nil {

		return sendResponse(c, errDecode.Error())
	}

	var generatedProducer Producer

	mapstructure.Decode(req, &generatedProducer)

	b.Queues[0].upsertProducer(&generatedProducer)

	return sendResponse(c, "ok")
}

func (b *Broker) handlerPostEvent(c *fiber.Ctx) error {
	req, errDecode := decodeRequest[requestAddEvent](c.Request())
	if errDecode != nil {
		return sendResponse(c, errDecode.Error())
	}

	var payload requestAddEvent

	mapstructure.Decode(req, &payload)

	if _, errGet := b.Queues[0].getProducer(payload.ProducerCode); errGet != nil {
		return sendResponse(c, errGet.Error())
	}

	b.Queues[0].AddMessage(
		&Message{
			Body: []byte(strings.Join(payload.Payload, "")),
		},
	)

	return sendResponse(c, "ok")
}

func (b *Broker) handlerConsumeEvent(c *fiber.Ctx) error {
	req, errDecode := decodeRequest[requestConsumeEvent](c.Request())
	if errDecode != nil {
		return sendResponse(c, errDecode.Error())
	}

	var consumer requestConsumeEvent

	mapstructure.Decode(req, &consumer)

	if _, errGet := b.Queues[0].getConsumer(consumer.ConsumerCode); errGet != nil {
		return sendResponse(c, errGet.Error())
	}

	var reconstructedEvent Message

	// TODO: get event from queue

	return sendResponse(c, string(reconstructedEvent.Body))
}
