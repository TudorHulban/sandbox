package main

const _Port = "8080"
const _SecondsDefaultQueueTTL = 600

const (
	urlRegisterProducer = "/registerproducer"
	urlRegisterConsumer = "/registerconsumer"
	urlPostEvents       = "/postevents"
	urlConsumeEvent     = "/consumevent" // switch to GET - later
	urlReadyEventID     = "/readyevent"
	urlStatusRequestID  = "/statusrequest" // switch to GET - later)
)
