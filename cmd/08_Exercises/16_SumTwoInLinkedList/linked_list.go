package main

import (
	"errors"
	"strconv"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList(n *Node) *LinkedList {
	return &LinkedList{
		Head: n,
	}
}

func NewLinkedListFromNumber(n int) *LinkedList {
	l := NewLinkedList(
		&Node{
			Value: n,
		},
	)

	stringRepresentation := strconv.Itoa(n)

	for i := len(stringRepresentation) - 2; i >= 0; i-- {
		numericDigit, _ := strconv.Atoi(stringRepresentation[i : i+1])

		l.Prepend(
			&Node{
				Value: numericDigit,
			},
		)
	}

	return l
}

func (l *LinkedList) Prepend(n *Node) {
	n.Next = l.Head

	l.Head = n
}

func (l *LinkedList) GetNumber() (int, error) {
	if l.Head == nil {
		return 0,
			errors.New("nil head value")
	}

	integers := []int{l.Head.Value}
	next := l.Head

	for next.Next != nil {
		integers = append(integers, next.Next.Value)

		next = next.Next
	}

	var result string

	for j := len(integers) - 1; j >= 0; j-- {
		result = strconv.Itoa(integers[j]) + result
	}

	return strconv.Atoi(result)
}
