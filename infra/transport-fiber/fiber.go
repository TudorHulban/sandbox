package transportfiber

import (
	"fmt"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type Transport struct {
	app *fiber.App

	Host string
	Port uint16
}

type ParamTransportSocket struct {
	Port string `valid:"port"`
	Host string `valid:"host"`
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

			app: fiber.New(),
		},
		nil
}

type ParamAddRoute struct {
	Handler func(*fiber.Ctx) error

	Method string
	Path   string
}

func (t *Transport) AddRoute(params *ParamAddRoute) {
	t.app.Add(
		params.Method,
		params.Path,
		params.Handler,
	)
}

func (t *Transport) getListeningSocket() string {
	return fmt.Sprintf(
		"%s:%d",
		t.Host,
		t.Port,
	)
}

func (t *Transport) Start() error {
	if len(t.app.GetRoutes()) == 0 {
		return fmt.Errorf(
			"no routes for %s transport",
			_nameTransport,
		)
	}

	fmt.Printf(
		"transport fiber started on http://%s",
		t.getListeningSocket(),
	)

	return t.app.Listen(
		t.getListeningSocket(),
	)
}
