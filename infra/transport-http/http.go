package transporthttp

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/asaskevich/govalidator"
)

type Transport struct {
	Port uint16
	Host string

	server *http.Server
}

type ParamTransportSocket struct {
	Port string `valid:"port"`
	Host string `valid:"host"`

	Handlers http.Handler
}

func NewTransport(params *ParamTransportSocket) (*Transport, error) {
	if _, errValidate := govalidator.ValidateStruct(params); errValidate != nil {
		return nil,
			errValidate
	}

	numericPort, errConversion := strconv.ParseUint(params.Port, 10, 64)
	if errConversion != nil {
		return nil,
			errConversion
	}

	return &Transport{
			Port: uint16(numericPort),
			Host: params.Host,

			server: &http.Server{
				Handler: params.Handlers,
			},
		},
		nil
}

func (t *Transport) getListeningSocket() string {
	return fmt.Sprintf(
		"%s:%d",
		t.Host,
		t.Port,
	)
}

func (t *Transport) Start(ctx context.Context) error {
	fmt.Printf(
		"transport %s started on http://%s",
		_nameTransport,
		t.getListeningSocket(),
	)

	chOSSignals := make(chan os.Signal, 1)
	signal.Notify(chOSSignals,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		select {
		case <-chOSSignals:
		case <-ctx.Done():
		}

		errShutdown := t.server.Shutdown(ctx)
		if errShutdown != nil {
			log.Fatal(errShutdown)
		}
	}()

	t.server.Addr = t.getListeningSocket()

	if errListenAndServe := t.server.ListenAndServe(); errListenAndServe != nil {
		return errListenAndServe
	}

	return nil
}
