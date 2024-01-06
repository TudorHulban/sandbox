package main

import "net/http"

func handleRoutes(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/login" {
		handlerLogin(w, r)

		return
	}

	if r.URL.Path == "/array" {
		handlerSignals(w, r)

		return
	}

	http.ServeFile(
		w,
		r,
		r.URL.Path[1:],
	)
}
