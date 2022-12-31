package main

import (
	"fmt"
	"runtime/debug"
)

type node struct {
	value int
	left  *node
	right *node
}

var (
	ExampleTree = &node{
		value: 1,
		left: &node{
			value: 2,
			left: &node{
				value: 3,
			},
			right: &node{
				value: 4,
			},
		},
		right: &node{
			value: 5,
		},
	}
)

func main() {
	fmt.Println(sumIterative(ExampleTree))
	fmt.Println(sumRecursive(ExampleTree))
	fmt.Println(max(ExampleTree))
	fmt.Println(MaxInline(ExampleTree))

	debug.SetMaxStack(262144000 * 2) // twice the 32-bit limit
	infiniteCount(0)
}

func infiniteCount(i int) {
	if i%1000 == 0 {
		fmt.Println(i)
	}
	infiniteCount(i + 1)
}

func sumRecursive(node *node) int {
	if node == nil {
		return 0
	}
	return node.value + sumRecursive(node.left) + sumRecursive(node.right)
}

func sumIterative(root *node) int {
	// not an ideal queue, but acts like a queue
	queue := make(chan *node, 10)
	queue <- root
	var sum int
	for {
		select {
		case node := <-queue:
			sum += node.value
			if node.left != nil {
				queue <- node.left
			}
			if node.right != nil {
				queue <- node.right
			}
		default:
			return sum
		}
	}
	return sum
}

func max(root *node) int {
	currentMax := 0

	var inner func(node *node)
	inner = func(node *node) {
		if node == nil {
			return
		}
		if node.value > currentMax {
			currentMax = node.value
		}
		inner(node.left)
		inner(node.right)
	}

	inner(root)

	return currentMax
}

func MaxInline(root *node) int {
	return maxInline(root, 0)
}

func maxInline(node *node, maxValue int) int {
	if node == nil {
		return maxValue
	}

	if node.value > maxValue {
		maxValue = node.value
	}

	maxLeft := maxInline(node.left, maxValue)
	maxRight := maxInline(node.right, maxValue)
	if maxLeft > maxRight {
		return maxLeft
	}
	return maxRight
}
