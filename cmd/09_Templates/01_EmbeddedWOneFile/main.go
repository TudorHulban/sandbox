package main

import (
	"embed"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

//go:embed templates/*.gohtml
var f embed.FS

func main() {
	s := server{
		app:  fiber.New(),
		port: ":3000",
	}

	state := State{
		Session: Session{
			IsAuthenticated: true,
			Username:        "john",
		},
		Data: []string{"Orange", "Apple", "Banana"},
	}

	s.app.Get("/",
		func(c *fiber.Ctx) error {
			buf := s.renderState(&state)

			return c.Type("html").Send(buf.Bytes())
		},
	)

	fmt.Printf(
		"listening on http://localhost%s\n",
		s.port,
	)

	log.Fatal(
		s.app.Listen(s.port),
	)
}
