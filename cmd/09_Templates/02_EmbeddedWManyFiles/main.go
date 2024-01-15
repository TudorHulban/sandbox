package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name    string
	IsAdmin bool
}

type server struct {
	app  *fiber.App
	port string
}

//go:embed templates/*.gohtml
var f embed.FS

func main() {
	s := server{
		app:  fiber.New(),
		port: ":3000",
	}

	user := User{
		IsAdmin: true,
		Name:    "john"}

	s.app.Get("/", func(c *fiber.Ctx) error {
		buf := s.renderContent(&user)

		return c.Type("html").Send(buf.Bytes())
	})

	log.Fatal(s.app.Listen(s.port))
}

func (s *server) renderContent(user *User) bytes.Buffer {
	funcMap := template.FuncMap{
		"isAdmin": func(isAdmin bool) bool {
			return isAdmin
		},
	}

	tmpl := template.New("views")
	tmpl.Funcs(funcMap)

	tmpl, err := tmpl.ParseFS(f, "templates/*.gohtml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var buf bytes.Buffer

	tmpl.ExecuteTemplate(&buf, "layout.gohtml", user)

	return buf
}
