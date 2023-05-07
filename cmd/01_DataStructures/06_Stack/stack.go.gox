package main

import (
	"errors"
	"strings"
)

type stack []any

func (s *stack) push(item any) {
	*s = append(*s, item)
}

func (s *stack) peek() any {
	if len(*s) == 0 {
		return nil
	}

	return (*s)[len(*s)-1]
}

func (s *stack) pop() (any, error) {
	if len(*s) == 0 {
		return "", errors.New("stack is empty")
	}

	res := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]

	return res, nil
}

func (s stack) String() string {
	var res []string

	for _, item := range s {
		res = append(res, item.(string))
	}

	return strings.Join(res, "/")
}
