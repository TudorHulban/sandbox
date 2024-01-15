package main

type Session struct {
	Username        string
	IsAuthenticated bool
}

type State struct {
	Session

	Data []string
}
