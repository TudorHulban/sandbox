package main

import (
	"strconv"
	"strings"
)

type stack []int

func (s *stack) push(element int) {
	newStack := []int{element}

	*s = append(newStack,
		*s...,
	)
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) peek() int {
	if s.isEmpty() {
		return 0
	}

	return (*s)[0]
}

func (s *stack) pop() int {
	if len(*s) == 0 {
		return 0
	}

	res := (*s)[0]

	*s = (*s)[1:]

	return res
}

func (s *stack) reverse() {
	frontElement := s.pop()

	if len(*s) == 0 {
		s.push(frontElement)

		return
	}

	s.reverse()

	*s = append(*s, frontElement)
}

func (s stack) String() string {
	result := make([]string, 0)

	for _, el := range []int(s) {
		result = append(result,
			strconv.Itoa(el),
		)
	}

	return strings.Join(result, "")
}
