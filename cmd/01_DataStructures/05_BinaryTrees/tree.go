package main

import (
	"fmt"
	"strings"
)

type node struct {
	leftNode  *node
	rightNode *node

	value int
}

func newNode(value int) *node {
	return &node{
		value: value,
	}
}

func (n *node) IsZero() bool {
	return (n.value == 0) && (n.leftNode == nil) && (n.rightNode == nil)
}

func (n *node) Add(key int) {
	if key > n.value {
		if n.rightNode == nil {
			n.rightNode = &node{
				value: key,
			}

			return
		}

		n.rightNode.Add(key)
	} else if key < n.value {
		if n.leftNode == nil {
			n.leftNode = &node{
				value: key,
			}

			return
		}

		n.leftNode.Add(key)
	}
}

func (n *node) HasKey(key int) bool {
	if n == nil {
		return false
	}

	if n.value == key {
		return true
	}

	if key > n.value {
		return n.rightNode.HasKey(key)
	}

	return n.leftNode.HasKey(key)
}

func (n *node) GetNode(withKey int, result *node) {
	if n == nil {
		return
	}

	if n.value == withKey {
		*result = *n

		return
	}

	if withKey > n.value {
		n.rightNode.GetNode(withKey, result)

		return
	}

	n.leftNode.GetNode(withKey, result)
}

func (n *node) String() string {
	res := []string{"tree data:"}

	chValues := make(chan int)
	defer close(chValues)

	chStop := make(chan struct{})
	defer close(chStop)

	go func() {
		var i int

		for {
			select {
			case value := <-chValues:
				res = append(res, fmt.Sprintf("%d: %d", i, value))
				i++

			case <-chStop:
				return
			}
		}
	}()

	walk(n, chValues)

	chStop <- struct{}{}

	return strings.Join(res, "\n")
}

func (n *node) Info() string {
	if n.IsZero() {
		return "node is nil"
	}

	res := []string{"node data:"}

	res = append(res, fmt.Sprintf("node value: %d", n.value))

	if n.leftNode != nil {
		res = append(res, fmt.Sprintf("node left value: %d", n.leftNode.value))
	} else {
		res = append(res, "node left: nil")
	}

	if n.rightNode != nil {
		res = append(res, fmt.Sprintf("node right value: %d", n.rightNode.value))
	} else {
		res = append(res, "node right: nil")
	}

	return strings.Join(res, "\n")
}
