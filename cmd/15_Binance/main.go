package main

import (
	"fmt"
	"os"
)

func main() {
	cfg := Config{
		URI: urlBinance,
	}

	c, errNew := NewClient(cfg)
	if errNew != nil {
		fmt.Println(errNew)

		os.Exit(1)
	}

	go c.ReadMessages()

	<-c.stop

	close(c.stop)
	close(c.interrupt)
}
