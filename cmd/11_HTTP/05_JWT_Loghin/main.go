package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HideBanner = true

	e.Use(
		middleware.CORSWithConfig(
			middleware.CORSConfig{
				AllowOrigins: []string{"*"},
				AllowMethods: []string{
					http.MethodGet,
					http.MethodPut,
					http.MethodPost,
					http.MethodDelete,
				},
			},
		),
	)

	//public routes
	e.Static("/", "../../public")
	e.POST("/login", handleLoginReq)
	e.GET("/", handlePublicRoutes)

	//restricted routes
	r := e.Group("/a")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", handleRestrictedRoutes)

	e.Logger.Fatal(e.Start(":3000"))
}

func handleLoginReq(c echo.Context) error {
	u := c.FormValue("username")
	p := c.FormValue("password")

	userData, errValidate := ValidateCredentials(u, p)
	if errValidate != nil {
		return echo.ErrInternalServerError
	}

	if userData.ID == -1 {
		return echo.ErrUnauthorized
	}

	token, errCrToken := createToken(&paramsCreateToken{
		fullName:         userData.FirstName + " " + userData.LastName,
		signingWith:      _Secret,
		isAdmin:          userData.IsAdmin,
		timeIntervalSecs: _SecondsTTLToken,
	})
	if errCrToken != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, token)
}
