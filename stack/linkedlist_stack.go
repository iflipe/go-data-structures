package stack

import (
	"errors"
	"main/list"
)

type LinkedListStack struct {
	head *list.Node
	size int
}

func (stack *LinkedListStack) Push(val int) {
	newNode := &list.Node{Value: val, Next: stack.head}
	stack.head = newNode
	stack.size++
}

func (stack *LinkedListStack) Pop() (int, error) {
	if !stack.Empty() {
		val := stack.head.Value
		stack.head = stack.head.Next
		return val, nil
	} else {
		return -1, errors.New("empty stack")
	}
}

func (stack *LinkedListStack) Peek() (int, error) {
	if !stack.Empty() {
		return stack.head.Value, nil
	} else {
		return -1, errors.New("empty stack")
	}
}

func (stack *LinkedListStack) Empty() bool {
	return stack.size == 0
}

func (stack *LinkedListStack) Size() int {
	return stack.size
}
