package main

import (
	"fmt"
	"time"
)

func main() {
	chDo := make(chan int64)
	defer close(chDo)

	chStop := make(chan struct{})
	defer close(chStop)

	go do(chDo)
	go stop(chStop)

loop:
	for {
		select {
		case <-chStop:
			break loop

		case msg := <-chDo:
			fmt.Println("fast work", msg)

		default:
			fmt.Println("slow work", time.Now().UnixMilli())
			time.Sleep(1500 * time.Millisecond)
		}
	}

	fmt.Println("exiting ....")
}
