package main

import (
	"fmt"
	"strings"

	"github.com/asaskevich/govalidator"
)

func NewURL(parts ...string) (string, error) {
	url := "http://" + strings.Join(parts, "/")

	if !govalidator.IsURL(url) {
		return "",
			fmt.Errorf("parts '%s' are not a URL", url)
	}

	return url, nil
}
