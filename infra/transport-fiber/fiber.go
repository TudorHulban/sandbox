package transportfiber

import "github.com/gofiber/fiber/v2"

type Transport struct {
	Port uint16

	app *fiber.App
}

type ParamTransportFiber struct {
	Port string
}

func NewTransport()
