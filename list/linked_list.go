package list

import (
	"errors"
)

type LinkedList struct {
	head     *Node
	inserted int
}

type Node struct {
	Value int
	Next  *Node
}

func (list *LinkedList) Add(value int) {
	newNode := &Node{value, nil}
	if list.head == nil {
		list.head = newNode
	} else {
		iter := list.head
		for iter.Next != nil {
			iter = iter.Next
		}
		iter.Next = newNode
	}
	list.inserted++
}

func (list *LinkedList) AddOnIndex(val int, index int) error {
	if index < 0 || index > list.inserted {
		return errors.New("index out of bounds")
	}

	newNode := &Node{val, nil}
	if list.inserted == 0 {
		list.head = newNode
	} else if index == 0 {
		newNode.Next = list.head.Next
		list.head = newNode
	} else {
		prev := list.head
		for i := 0; i < index-1; i++ {
			prev = prev.Next
		}
		newNode.Next = prev.Next
		prev.Next = newNode
	}
	list.inserted++
	return nil
}

func (list *LinkedList) RemoveOnIndex(index int) error {
	if index < 0 || index >= list.inserted {
		return errors.New("index out of bounds")
	}
	if index == 0 {
		list.head = list.head.Next
	} else {
		iter := list.head
		for i := 0; i < index-1; i++ {
			iter = iter.Next
		}
		iter.Next = iter.Next.Next
	}
	list.inserted--
	return nil
}

func (list *LinkedList) Get(index int) (int, error) {
	if index < 0 || index >= list.inserted {
		return index, errors.New("index out of bounds")
	}

	iter := list.head
	for i := 0; i < index; i++ {
		iter = iter.Next
	}
	return iter.Value, nil
}

func (list *LinkedList) Set(value int, index int) error {
	if index < 0 || index >= list.inserted {
		return errors.New("index out of bounds")
	}
	iter := list.head
	for i := 0; i < index; i++ {
		iter = iter.Next
	}
	iter.Value = value
	return nil
}

func (list *LinkedList) Size() int {
	return list.inserted
}
