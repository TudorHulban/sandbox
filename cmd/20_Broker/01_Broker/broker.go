package main

import (
	"net/http"

	transportfiber "github.com/TudorHulban/GolangSandbox/infra/transport-fiber"
)

type Broker struct {
	Transport *transportfiber.Transport
	Queues    []*Queue
}

type ParamTransport struct {
	Port string
	Host string
}

func NewBroker(params *ParamTransport) (*Broker, error) {
	transport, errCrTransport := transportfiber.NewTransport(
		(*transportfiber.ParamTransportSocket)(params),
	)
	if errCrTransport != nil {
		return nil,
			errCrTransport
	}

	return &Broker{
			Transport: transport,
			Queues: []*Queue{
				newQueue(_SecondsDefaultQueueTTL),
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
