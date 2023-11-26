package main

import (
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

func (s *stack) pop() any {
	if len(*s) == 0 {
		return nil
	}

	res := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]

	return res
}

func (s stack) String() string {
	var res []string

	for _, item := range s {
		res = append(res, item.(string))
	}

	return strings.Join(res, "/")
}
