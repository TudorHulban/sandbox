package main

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type paramsCreateToken struct {
	fullName         string
	signingWith      string
	timeIntervalSecs int
	isAdmin          bool
}

func createToken(params *paramsCreateToken) (map[string]string, error) {
	jwtToken := jwt.New(jwt.SigningMethodHS256)

	claims := jwtToken.Claims.(jwt.MapClaims)
	claims["name"] = params.fullName
	claims["admin"] = strconv.FormatBool(params.isAdmin)
	claims["exp"] = time.Now().
		Add(time.Second * time.Duration(params.timeIntervalSecs)).Unix()

	signedToken, err := jwtToken.SignedString([]byte(params.signingWith))
	if err != nil {
		return nil, err
	}

	return map[string]string{"token": signedToken}, nil
}
