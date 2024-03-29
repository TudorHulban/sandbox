package main

import (
	"errors"
	"fmt"
	"strings"
)

type node struct {
	nextNode *node

	data int
}

type linkedList struct {
	head *node
}

func newLinkedListFromHead(head *node) *linkedList {
	return &linkedList{
		head: head,
	}
}

func newLinkedListFromNodes(nodes ...*node) *linkedList {
	if len(nodes) == 0 {
		return nil
	}

	res := newLinkedListFromHead(nodes[0])

	res.prepend(nodes[1:]...)

	return res
}

func (l *linkedList) prepend(nodes ...*node) {
	if len(nodes) == 0 {
		return
	}

	for _, node := range nodes {
		node.nextNode = l.head

		l.head = node
	}
}

func (l *linkedList) reverse() {
	currentNode := l.head
	nextNode := l.head.nextNode

	var wasFirst bool

	for nextNode != nil {
		l.head = nextNode

		nextNode = l.head.nextNode
		l.head.nextNode = currentNode

		if !wasFirst {
			currentNode.nextNode = nil
		}

		wasFirst = true

		currentNode = l.head
	}
}

func (l *linkedList) String() string {
	var result []string

	nextNode := l.head

	for nextNode != nil {
		result = append(result, fmt.Sprintf("%d", nextNode.data))

		nextNode = nextNode.nextNode
	}

	return strings.Join(result, "")
}

func (l *linkedList) deleteValue(valueData int) error {
	if l.head == nil {
		return errors.New("empty list")
	}

	if l.head.data == valueData {
		if l.head.nextNode != nil {
			l.head = l.head.nextNode

			return nil
		}

		l.head = nil

		return errors.New("empty list")
	}

	// if previous did not trigger the one value, list does not have the value.
	if l.head.nextNode == nil {
		return fmt.Errorf(errorNotFound, valueData)
	}

	next := l.head

	for next.nextNode.data != valueData {
		if next.nextNode.nextNode == nil {
			return fmt.Errorf(errorNotFound, valueData)
		}

		if next.nextNode == nil {
			return errors.New("stopping , only head remains")
		}

		next = next.nextNode
	}

	next.nextNode = next.nextNode.nextNode

	return nil
}
