package transportfiber

import (
	"fmt"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type Transport struct {
	Port uint16
	Host string

	app *fiber.App
}

type ParamTransportFiber struct {
	Port string
	Host string
}

func NewTransport(params *ParamTransportFiber) (*Transport, error) {
	if !govalidator.IsHost(params.Port) {
		return nil,
			fmt.Errorf(
				"passed port: %s is not valid",
				params.Port)
	}

	if !govalidator.IsPort(params.Port) {
		return nil,
			fmt.Errorf(
				"passed port: %s is not valid",
				params.Port)
	}

	numericPort, errConversion := strconv.ParseUint(params.Port, 10, 10)
	if errConversion != nil {
		return nil,
			errConversion
	}

	return &Transport{
			Port: uint16(numericPort),
			app:  fiber.New(),
		},
		nil
}

type ParamAddRoute struct {
	Method  string
	Path    string
	Handler func(*fiber.Ctx) error
}

func (t *Transport) AddRoute(params *ParamAddRoute) {
	t.app.Add(
		params.Method,
		params.Path,
		params.Handler,
	)
}

func (t *Transport) Start() error {
	if len(t.app.GetRoutes()) == 0 {
		return fmt.Errorf(
			"no routes for %s transport",
			_nameTransport,
		)
	}

	return t.app.Listen(
		fmt.Sprintf(
			"%s:%d",
			t.Host,
			t.Port,
		),
	)
}
