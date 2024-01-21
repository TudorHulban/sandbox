package app

import "github.com/gofiber/fiber/v2"

func (app *App) HandlerAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("alive")
	}
}
