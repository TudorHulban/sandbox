package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	chDo := make(chan string)
	defer close(chDo)

	chStop := make(chan struct{})
	defer close(chStop)

	go stop(chStop)
	go do(chDo)

loop:
	for {
		select {
		case <-chStop:
			{
				break loop
			}

		case msg := <-chDo:
			{
				fmt.Println(msg)
			}

		default:
			{
				fmt.Println("other work")
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}

	fmt.Println("exiting ....")
}

func do(c chan string) {
	for {
		c <- strconv.FormatInt(time.Now().Unix(), 10)

		time.Sleep(1000 * time.Millisecond)
	}
}

func stop(c chan struct{}) {
	time.Sleep(3000 * time.Millisecond)

	c <- struct{}{}
}
