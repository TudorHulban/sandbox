package main

type Producer struct {
	Code string `json:"code"`
}

type Consumer struct {
	Code   string `json:"code"`
	Socket string `json:"socket"`
}

type requestAddEvent struct {
	ProducerCode string   `json:"code"`
	ID           int64    `json:"id"`
	SecondsTTL   int64    `json:"ttl"`
	Payload      []string `json:"payload"`
}

type readyevent struct {
	ConsumerCode string `json:"code"`
	EventID      int64  `json:"id"`
}

type requeststatus struct {
	ProducerCode string `json:"code"`
	RequestID    int64  `json:"id"`
}

type requestConsumeEvent struct {
	ConsumerCode string `json:"code"`
}
