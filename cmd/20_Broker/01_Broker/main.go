package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mitchellh/mapstructure"
)

type jsonAnswer struct {
	Response string `json:"response"`
}

func main() {
	http.HandleFunc(urlRegisterConsumer, handlerRegisterConsumer)
	http.HandleFunc(urlPostEvents, handlerPostEvent)
	http.HandleFunc(urlConsumeEvent, handlerConsumeEvent)
	http.HandleFunc(urlReadyEventID, handlerReadyEvID)
	http.HandleFunc(urlStatusRequestID, handlerStatusRequestID)

	http.ListenAndServe((fmt.Sprintf(":%v", _Port)), nil)
}

func handlerRegisterConsumer(w http.ResponseWriter, r *http.Request) {
	req, errDecode := decodeRequest(r, reqisterConsumer)
	if errDecode != nil {
		sendResponse(w, errDecode.Error())

		return
	}

	var generatedConsumer consumer
	mapstructure.Decode(req, &generatedConsumer)
	q.addConsumer(&generatedConsumer)

	fmt.Println("registered consumers:", len(q.consumers), q.consumers)

	sendResponse(w, "ok")
}

func handlerPostEvent(w http.ResponseWriter, r *http.Request) {
	log.Println("-- post event")

	req, errDecode := decodeRequest(r, requestEvent)
	if errDecode != nil {
		log.Println("event decode error: ", errDecode)

		sendResponse(w, errDecode.Error())

		return
	}

	var payload request
	mapstructure.Decode(req, &payload)

	if exists, _ := q.isRegisteredProducer(payload.ProducerCode); !exists {
		sendResponse(w, "not registered")

		return
	}

	q.addPayload(payload)

	log.Println("registered events:", len(q.events))
	log.Println(q.events)

	sendResponse(w, "ok")
}

func handlerConsumeEvent(w http.ResponseWriter, r *http.Request) {
	req, errDecode := decodeRequest(r, requestTask)
	if errDecode != nil {
		sendResponse(w, errDecode.Error())

		return
	}

	var generatedTask task
	mapstructure.Decode(req, &generatedTask)
	log.Println("requesting work:", generatedTask.ConsumerCode)

	exists, consumer := q.isRegisteredConsumer(generatedTask.ConsumerCode)
	if !exists {
		sendResponse(w, "not registered")

		return
	}

	event := q.getOldestEvent(consumer)
	if event == nil {
		sendResponse(w, "no work")

		return
	}

	log.Println("consumed event:", event)
	log.Println(q.events)

	sendResponse(w, event.request)
}

// handlerReadyEvID - invoked by consumer
func handlerReadyEvID(w http.ResponseWriter, r *http.Request) {
	req, errDecode := decodeRequest(r, readyEvent)
	if errDecode != nil {
		sendResponse(w, errDecode.Error())

		return
	}

	var generatedEvent readyevent
	mapstructure.Decode(req, &generatedEvent)

	log.Println("ready event:", generatedEvent.EventID)

	event := q.getEventByID(generatedEvent.EventID)
	if event == nil {
		sendResponse(w, "event does not exist in queue")

		return
	}

	if event.consumedByCode != generatedEvent.ConsumerCode {
		sendResponse(w, "consumed wrong event id")

		return
	}

	event.consumeEnd = time.Now().UnixNano()
	log.Println("consumed event:", event)
	log.Println(q.events)

	sendResponse(w, event.request)
}

// handlerStatusRequestID - invoked by producer
func handlerStatusRequestID(w http.ResponseWriter, r *http.Request) {
	req, errDecode := decodeRequest(r, requestStatus)
	if errDecode != nil {
		log.Println("event decode error: ", errDecode)
		sendResponse(w, errDecode.Error())
		return
	}

	var generatedEvent readyevent
	mapstructure.Decode(req, &generatedEvent)

	log.Println("ready event:", generatedEvent.EventID)

	event := q.getEventByID(generatedEvent.EventID)
	if event == nil {
		sendResponse(w, "event does not exist in queue")

		return
	}

	if event.consumedByCode != generatedEvent.ConsumerCode {
		sendResponse(w, "consumed wrong event id")
		return
	}

	event.consumeEnd = time.Now().UnixNano()

	log.Println("consumed event:", event)
	log.Println(q.events)

	sendResponse(w, event.request)
}
