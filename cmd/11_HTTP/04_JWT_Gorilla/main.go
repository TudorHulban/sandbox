package main

//check also https://dzone.com/articles/authentication-in-golang-with-jwts

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const HTTPPort = 3000
const URLLogin = "/login"
const URLRestricted = "/a"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HandlerPublic)
	r.HandleFunc("/login", HandlerLogin)
	r.Handle("/a", authMiddleware(handlerAdmin))

	r.PathPrefix("/public").Handler(
		http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))),
	)

	log.Fatal(
		http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r)),
	)
}

func HandlerPublic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Public Route")
}

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login Route")
}

var handlerAdmin = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Admin Route")
	},
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if len(tokenString) == 0 {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte("Missing Authorization Header"))

				return
			}

			tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

			claims, errVerifyToken := verifyToken(tokenString)
			if errVerifyToken != nil {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte("verifying JWT token: " + errVerifyToken.Error()))

				return
			}

			name := claims.(jwt.MapClaims)["name"].(string)
			role := claims.(jwt.MapClaims)["role"].(string)

			r.Header.Set("name", name)
			r.Header.Set("role", role)

			next.ServeHTTP(w, r)
		},
	)
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte("keymaker")

	token, errParse := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (any, error) {
			return signingKey, nil
		},
	)
	if errParse != nil {
		return nil, errParse
	}

	return token.Claims, errParse
}
