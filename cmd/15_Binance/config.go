package main

import "net/http"

type Config struct {
	RequestHeader       http.Header
	URI                 string
	PongIntervalSeconds uint
}
