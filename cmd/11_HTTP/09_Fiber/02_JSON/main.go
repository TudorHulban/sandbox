package main

// test with curl http://127.0.0.1:3000

import "github.com/gofiber/fiber/v2"

type Person struct {
	Name string
	Age  int
}

type Team []Person

func main() {
	app := fiber.New()

	app.Get("/", serve)

	app.Listen(":3000")
}

func serve(c *fiber.Ctx) error {
	team := []Person{
		{
			Name: "John",
			Age:  34,
		},
		{
			Name: "Ann",
			Age:  32,
		},
	}

	return c.JSON(team)
}
