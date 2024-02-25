package transporthttp

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestHTTP(t *testing.T) {
	routerChi := chi.NewRouter()

	routerChi.Get(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("alive"))
		},
	)

	transportHTTP, errNew := NewTransport(
		&ParamTransportSocket{
			Port: "3001",
			Host: "localhost",

			Handlers: routerChi,
		},
	)
	require.NoError(t, errNew)
	require.NotZero(t, transportHTTP)

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*2500)
	defer cancel()

	transportHTTP.Start(ctx)
}
