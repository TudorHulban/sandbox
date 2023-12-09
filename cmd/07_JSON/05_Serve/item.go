package main

type item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type items struct {
	Items []item `json:items`
}

var _toServeItems items
