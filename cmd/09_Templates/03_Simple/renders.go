package main

import (
	"fmt"
	"io"
	"os"
	"text/template"
)

// f1 renders template to stdout.
func f1(w io.Writer, templateName string, model Person) {
	t := template.New(templateName)

	t, errParse := t.Parse("hi {{.Name}}!")
	if errParse != nil {
		fmt.Println("errParse: ", errParse)

		return
	}

	t.Execute(w, model)
}

// f2 loads template and renders it to stdout.
func f2(w io.Writer, templateFilePath string, model Person) {
	t, errParse := template.ParseFiles(templateFilePath)
	if errParse != nil {
		fmt.Println("errParse: ", errParse)
	}

	if errExec := t.Execute(w, model); errExec != nil {
		fmt.Println("errExec: ", errExec)
	}
}

// f3 Loads template and renders it to file.
func f3(w io.Writer, templateFilePath, renderToPath string, model Person) {
	t, errParse := template.ParseFiles(templateFilePath)
	if errParse != nil {
		fmt.Println("errParse: ", errParse)
	}

	file, errCreate := os.Create(renderToPath)
	if errCreate != nil {
		fmt.Println("errCreate: ", errCreate)
	}
	defer file.Close()

	if errExec := t.Execute(w, model); errExec != nil {
		fmt.Println("errExec: ", errExec)
	}
}
