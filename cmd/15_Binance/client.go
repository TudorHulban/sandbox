package main

import (
	"fmt"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

type Client struct {
	connection *websocket.Conn
	URL        url.URL

	stop      chan struct{}
	interrupt chan os.Signal
}

func NewClient(cfg Config) (*Client, error) {
	url, errParse := url.Parse(cfg.URI)
	if errParse != nil {
		return nil, errParse
	}

	conn, _, errConn := websocket.DefaultDialer.Dial(
		url.String(),
		cfg.RequestHeader,
	)
	if errConn != nil {
		return nil, errConn
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	return &Client{
		connection: conn,
		URL:        *url,
		stop:       make(chan struct{}),
		interrupt:  interrupt,
	}, nil
}

func (c *Client) ReadMessages() {
loop:
	for {
		select {
		case <-c.interrupt:
			{
				fmt.Println("interrupt")

				break loop
			}
		default:
			{
				_, message, errRead := c.connection.ReadMessage()
				if errRead != nil {
					fmt.Println("read glitch:", errRead)

					return
				}

				go fmt.Println(string(message))
			}
		}
	}

	c.stop <- struct{}{}
}
