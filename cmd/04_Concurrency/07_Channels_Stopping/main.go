package main

import (
	"fmt"
	"time"
)

func do(work chan int64) {
	for {
		work <- time.Now().UnixMilli()
	}
}

func stop(stop chan struct{}) {
	time.Sleep(5000 * time.Millisecond)

	fmt.Println("elapsed finished...")

	stop <- struct{}{}
}

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
