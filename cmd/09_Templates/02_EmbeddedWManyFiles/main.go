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
		app:   fiber.New(),
		users: NewUsers(),

		port: ":3000",
	}

	s.app.Get("/",
		func(c *fiber.Ctx) error {
			user, errGetUser := s.users.GetUser(
				Token(
					string(
						c.Request().Header.Peek(_keyAuthorization),
					),
				),
			)
			if errGetUser != nil {
				return errGetUser
			}

			buf := s.renderContent(user)

			return c.Type("html").Send(buf.Bytes())
		},
	)

	fmt.Printf(
		"listening on http://localhost%s\n",
		s.port,
	)

	log.Fatal(s.app.Listen(s.port))
}
