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

	result := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]

	return result
}

func (s stack) String() string {
	result := make([]string, len(s))

	for ix, item := range s {
		result[ix] = item.(string)
	}

	return strings.Join(result, "/")
}
