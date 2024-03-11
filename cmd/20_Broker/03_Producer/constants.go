package main

// TODO: move to configuration file read.

const ip = "127.0.0.1"
const producerCode = "P1"
const defaultSecondsEventTTL = 600
const brokerAPIurl = "http://localhost:8080"
const routeRegister = "/registerproducer"
const routeEvents = "/postevents"
const routeStatus = "/statusrequest"
