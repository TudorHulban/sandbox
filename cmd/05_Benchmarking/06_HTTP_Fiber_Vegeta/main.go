package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	app := fiber.New()

	app.Get("/",
		func(c *fiber.Ctx) error {
			return c.SendStatus(http.StatusOK)
		},
	)

	go func() {
		_ = app.Listen(":3000")
	}()

	chStop1 := make(chan struct{})
	defer close(chStop1)

	chStop2 := make(chan struct{})
	defer close(chStop2)

	go func() {
		<-chStop1

		fmt.Println("Gracefully shutting down Fiber server.")

		_ = app.Shutdown()

		chStop2 <- struct{}{}
	}()

	rate := vegeta.Rate{
		Freq: 10,
		Per:  time.Second,
	}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(
		vegeta.Target{
			Method: "GET",
			URL:    "http://localhost:3000/",
		},
	)

	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics

	for res := range attacker.Attack(
		targeter,
		rate,
		duration,
		"Fiber!",
	) {
		metrics.Add(res)
	}

	metrics.Close()

	chStop1 <- struct{}{}

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
	fmt.Printf("95th percentile: %s\n", metrics.Latencies.P95)
	fmt.Printf("50th percentile: %s\n", metrics.Latencies.P50)
	fmt.Printf("requests: %d\n", metrics.Requests)
	fmt.Printf("Metrics errors: %#v\n", metrics.Errors)

	<-chStop2
}
