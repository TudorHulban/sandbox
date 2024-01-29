package ddltable

import plurals "github.com/gertd/go-pluralize"

func pluralize(word string) string {
	pluralize := plurals.NewClient()

	return pluralize.Pluralize(word, 2, false)
}
