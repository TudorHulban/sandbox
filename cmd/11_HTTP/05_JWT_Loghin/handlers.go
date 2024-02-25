package main

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func handlePublicRoutes(c echo.Context) error {
	return c.String(http.StatusOK, "Public Route")
}

func handleRestrictedRoutes(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return c.String(
		http.StatusOK,
		"Welcome "+claims["name"].(string)+"!",
	)
}
