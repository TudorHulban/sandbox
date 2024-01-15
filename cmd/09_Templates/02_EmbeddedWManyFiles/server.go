package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	app   *fiber.App
	users *Users

	port string
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

	errTemplate := tmpl.ExecuteTemplate(
		&buf,
		"layout.gohtml",
		user,
	)
	if errTemplate != nil {
		fmt.Println(errTemplate)

		os.Exit(1)
	}

	return buf
}
