package stack

import "errors"

type ArrayStack struct {
	v   []int
	top int
}

func (stack *ArrayStack) doubleStack() {
	doubledValues := make([]int, 2*len(stack.v))
	for i := 0; i <= stack.top; i++ {
		doubledValues[i] = stack.v[i]
	}
	stack.v = doubledValues
}

func (stack *ArrayStack) Push(value int) {
	if stack.top == len(stack.v)-1 {
		stack.doubleStack()
	}
	stack.top++
	stack.v[stack.top] = value

}

func (stack *ArrayStack) Pop() (int, error) {
	if !stack.Empty() {
		stack.top--
		return stack.v[stack.top+1], nil
	}
	return -1, errors.New("Empty stack")
}

func (stack *ArrayStack) Peek() (int, error) {
	if stack.Empty() {
		return -1, errors.New("Empty stack")
	}

	return stack.v[stack.top], nil
}

func (stack *ArrayStack) Empty() bool {
	return stack.top == -1
}

func (stack *ArrayStack) Size() int {
	return stack.top + 1
}

func (stack *ArrayStack) Init(size int) {
	if size >= 0 {
		stack.v = make([]int, size)
		stack.top = -1
	}
}
