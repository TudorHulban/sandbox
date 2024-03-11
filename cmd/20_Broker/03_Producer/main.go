package main // needs refactoring with REST, Websockets and methods on request including marshalling

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	registered, errRegister := registerWithBroker()
	if errRegister != nil {
		fmt.Println(errRegister)

		os.Exit(1)
	}

	fmt.Println("broker: ", registered, errRegister)
}

func cleanup() {
	fmt.Println("cleanup")
}

func main() {
	chOS := make(chan os.Signal, 1)

	signal.Notify(chOS, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-chOS

		cleanup()

		os.Exit(1)
	}()

	var evposted bool

	for {
		if !evposted {
			event := request{
				id:           1,
				codeProducer: producerCode,
				ttl:          defaultSecondsEventTTL,
				payload: []string{
					"https://golang.org",
					"https://jsonformatter.curiousconcept.com",
				},
			}

			posted, errPost := postRequest(event)

			go getStatus(event)

			if errPost != nil {
				fmt.Println(errPost)

				time.Sleep(10 * time.Second)
			}

			fmt.Println("resp: ", posted)

			evposted = true
		}

		time.Sleep(10 * time.Second)

		fmt.Println("working...")
	}
}
