package main

import (
	"net/http"

	transportfiber "github.com/TudorHulban/GolangSandbox/infra/transport-fiber"
)

type Broker struct {
	Transport *transportfiber.Transport
	Queues    []*Topic
}

type ParamsNewBroker struct {
	transportfiber.ParamTransportSocket
}

func NewBroker(params *ParamsNewBroker) (*Broker, error) {
	transport, errCrTransport := transportfiber.NewTransport(
		&params.ParamTransportSocket,
	)
	if errCrTransport != nil {
		return nil,
			errCrTransport
	}

	return &Broker{
			Transport: transport,
			Queues: []*Topic{
				NewTopic(_SecondsDefaultQueueTTL),
			},
		},
		nil
}

func (b *Broker) AddRoutes() {
	b.Transport.AddRoute(
		&transportfiber.ParamAddRoute{
			Handler: b.handlerRegisterProducer,
			Method:  http.MethodGet,
			Path:    urlRegisterProducer,
		},
	)
}
