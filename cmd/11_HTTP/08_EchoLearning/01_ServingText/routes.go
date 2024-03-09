package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func defaultRouteHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Aloha, World!")
}

func loginRouteHandler(c echo.Context) error {
	valuesForm, err := c.FormParams()
	if err != nil {
		log.Fatal("loginRouteHandler:", err)
	}

	for key, value := range valuesForm {
		fmt.Println(key, value[0])
	}

	return err
}
