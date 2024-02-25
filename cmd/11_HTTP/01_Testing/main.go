package main

import (
	"fmt"
	"io"
	"net/http"
)

// Req: http://localhost:1234/upper?word=xxx
// Res: XXX

type client struct {
	url string
}

func newClient(url string) *client {
	return &client{
		url: url,
	}
}

func (c client) upperCase(word string) (string, error) {
	response, errGet := http.Get(c.url + "/upper" + word)
	if errGet != nil {
		return "", fmt.Errorf("GET error: %w", errGet)
	}
	defer response.Body.Close()

	body, errRead := io.ReadAll(response.Body)
	if errRead != nil {
		return "", fmt.Errorf("body error: %w", errGet)
	}

	return string(body), nil
}

func main() {}
