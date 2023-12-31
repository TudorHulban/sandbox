package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	chCancel := make(chan os.Signal, 1)

	signal.Notify(chCancel, os.Interrupt, syscall.SIGTERM)

	chStopScheduling := schedule(
		func() {
			log.Println("EVENT Triggered")
		},

		3*time.Second,
	)

	go func() {
		<-chCancel

		chStopScheduling <- struct{}{}
		close(chStopScheduling)

		cleanup()

		os.Exit(1)
	}()

	log.Println("going to sleep now...")

	for {
		time.Sleep(5 * time.Second) // or runtime.Gosched()

		log.Println("awaken... going back to sleep...")
	}
}
