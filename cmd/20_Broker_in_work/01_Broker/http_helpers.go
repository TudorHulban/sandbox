package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/valyala/fasthttp"
)

func decodeRequest[T any](r *fasthttp.Request) (*T, error) {
	decoder := json.NewDecoder(r.BodyStream())

	var result T

	if errDecode := decoder.Decode(&result); errDecode != nil {
		return nil,
			errDecode
	}

	return &result,
		nil
}

func readResponse(resp *http.Response) (string, error) {
	defer resp.Body.Close()

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return "", errRead
	}

	var result jsonAnswer

	errConvert := json.Unmarshal(body, &result)

	return result.Response, errConvert
}
