package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientUpperCase(t *testing.T) {
	word := "xxx"
	expectedResponse := "XXX"

	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "%s", expectedResponse)
			},
		),
	)
	defer server.Close()

	client := newClient(server.URL)
	response, errCall := client.upperCase(word)
	require.NoError(t, errCall)

	response = strings.TrimSpace(response)
	require.Equal(t, expectedResponse, response)
}
