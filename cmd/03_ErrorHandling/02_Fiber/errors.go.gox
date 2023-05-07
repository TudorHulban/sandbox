package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const errorMsgValidateInput = "validateInput: %w"

var errorBadInput = errors.New("bad input")

type errorBadText struct {
	badText string
}

var _ error = &errorBadText{}

// Error uses value receiver for simpler usage.
func (e errorBadText) Error() string {
	return fmt.Sprintf("bad text is: %s", e.badText)
}

type errorApp struct {
	Area string
	CODE string
	Err  error
}

func (e errorApp) Error() string {
	res := []string{
		"App Error:",
		fmt.Sprintf("Area:%s", e.Area),
		fmt.Sprintf("Code:%s", e.CODE),
		fmt.Sprintf("Error:%v", e.Err),
	}

	return strings.Join(res, "\n")
}

func (e errorApp) MarshalJSON() ([]byte, error) {
	res := make(map[string]any, 2)
	res["success"] = "false"

	resError := make(map[string]string, 3)
	resError["area"] = e.Area
	resError["code"] = e.CODE
	resError["error"] = e.Err.Error()

	res["error"] = resError

	return json.Marshal(&res)
}
