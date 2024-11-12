package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	process(1, 2, 3, 4)
}

func process(payloads ...int) {
	chReady := make(chan int)
	chCancel := make(chan os.Signal, 1)

	signal.Notify(chCancel, os.Interrupt, syscall.SIGTERM)

	for payload := range payloads {
		go func() {
			buf := payload

			time.Sleep(5 * time.Second)

			chReady <- buf
		}()

		select {
		case number := <-chReady:
			fmt.Println(number)

		case <-chCancel:
			fmt.Println("os cancel received. exiting ...")

			return
		}
	}
}
