package main

import (
	"fmt"
	"time"
)

func do(c chan int64) {
	for {
		c <- time.Now().UnixMilli()
	}
}

func stop(c chan struct{}) {
	time.Sleep(5000 * time.Millisecond)

	fmt.Println("elapsed finished...")

	c <- struct{}{}
}
