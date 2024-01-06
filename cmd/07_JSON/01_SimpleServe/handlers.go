package main

import (
	"encoding/json"
	"net/http"
)

// handlerLogin Handler with no marshaling, just returning slice of bytes.
func handlerLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, _ = w.Write(
		[]byte(
			`[{"value":101},{"value":78}]`,
		),
	)
}

func handlerSignals(w http.ResponseWriter, r *http.Request) {
	marshalledSignals, errMa := json.Marshal(state)
	if errMa != nil {
		_, _ = w.Write([]byte(errMa.Error()))

		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, _ = w.Write(marshalledSignals)
}
