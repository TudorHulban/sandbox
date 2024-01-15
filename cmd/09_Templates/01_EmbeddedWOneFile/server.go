package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	app *fiber.App

	port string
}

func (s *server) renderState(state *State) bytes.Buffer {
	funcMap := template.FuncMap{
		"isChosen": func(fruit string) bool {
			return fruit == _chosenFruit
		},
	}

	tmpl := template.New("")
	tmpl.Funcs(funcMap)

	tmpl, errParse := tmpl.ParseFS(f, "templates/index.gohtml")
	if errParse != nil {
		fmt.Println(errParse)

		os.Exit(1)
	}

	var buf bytes.Buffer

	errTemplate := tmpl.ExecuteTemplate(
		&buf,
		"index.gohtml",
		state,
	)
	if errTemplate != nil {
		fmt.Println(errTemplate)

		os.Exit(1)
	}

	return buf
}
