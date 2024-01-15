package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name    string
	IsAdmin bool
}

var userJohn = User{
	IsAdmin: true,
	Name:    "John",
}

var userMary = User{
	IsAdmin: false,
	Name:    "Mary",
}

var userTom = User{
	IsAdmin: false,
	Name:    "Tom",
}

type Token string

type Users struct {
	users map[Token]*User

	mu sync.Mutex
}

func NewUsers() *Users {
	return &Users{
		users: map[Token]*User{
			"1234": &userJohn,
			"2345": &userMary,
			"4567": &userTom,
		},
	}
}

func (u *Users) GetUser(token Token) (*User, error) {
	var result *User

	var exists bool

	u.mu.Lock()
	defer u.mu.Unlock()

	if result, exists = u.users[token]; !exists {
		return nil,
			fmt.Errorf("user for token %s not found",
				token,
			)
	}

	return result, nil
}
