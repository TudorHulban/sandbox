package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
)

func (b *Broker) handlerRegisterProducer(c *fiber.Ctx) error {
	req, errDecode := decodeRequest(c.Request(), registerProducer)
	if errDecode != nil {

		return sendResponse(c, errDecode.Error())
	}

	var generatedProducer producer
	mapstructure.Decode(req, &generatedProducer)

	b.Queues[0].addProducer(&generatedProducer)

	return sendResponse(c, "ok")
}
