package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get(
		"/:p",

		echo,
	)

	app.Listen(":7000")
}

func echo(c *fiber.Ctx) error {
	return c.SendString(c.Params("p"))
}
