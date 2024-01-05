package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Millisecond*500,
	)
	defer cancel()

	myfunc(ctx, time.Millisecond*600, "John")
}

func myfunc(ctx context.Context, d time.Duration, arg string) string {
	select {
	case <-ctx.Done():
		log.Print(ctx.Err())

	case <-time.After(d):
		fmt.Printf(
			"Hi %s\n",
			arg,
		) //some work: time.Sleep(time.Millisecond * 1000)
	}

	return "end"
}
