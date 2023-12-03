package main

func walk(n *node, chValues chan int) {
	if n == nil {
		return
	}

	chValues <- n.value

	if n.rightNode != nil {
		walk(n.rightNode, chValues)
	}

	if n.leftNode != nil {
		walk(n.leftNode, chValues)
	}
}
