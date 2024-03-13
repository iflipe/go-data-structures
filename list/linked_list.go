package list

import (
	"errors"
)

type LinkedList struct {
	head     *Node
	inserted int
}

type Node struct {
	value int
	next  *Node
}

func (list *LinkedList) Add(value int) {
	newNode := &Node{value, nil}
	if list.head == nil {
		list.head = newNode
	} else {
		iter := list.head
		for iter != nil {

			iter = iter.next
		}
		iter = newNode
	}
	list.inserted++
}

func (list *LinkedList) AddOnIndex(val int, index int) error {
	if index >= 0 && index <= list.inserted {
		newNode := &Node{val, nil}
		if list.head == nil {
			list.head = newNode
		} else {
			prev := list.head
			for i := 0; i < index-1; i++ {
				prev = prev.next
			}
			newNode.next = prev.next
			prev.next = newNode
		}
		list.inserted++
		return nil
	} else {
		return errors.New("error msg")
	}
}

func (list *LinkedList) RemoveOnIndex(index int) error {
	if index < 0 || index >= list.inserted {
		return errors.New("index out of bounds")
	}
	iter := list.head
	for i := 0; i < index-1; i++ {
		iter = iter.next
	}
	iter.next = iter.next.next
	list.inserted--
	return nil
}

func (list *LinkedList) Get(index int) (int, error) {
	if index >= 0 && index < list.inserted {
		iter := list.head
		for i := 0; i < index; i++ {
			iter = iter.next
		}
		return (iter.value), nil
	} else {
		return index, errors.New("index out of bounds")
	}
}

func (list *LinkedList) Set(value int, index int) error {
	if index < 0 && index >= list.inserted {
		return errors.New("index out of bounds")
	}
	iter := list.head
	for i := 0; i < index; i++ {
		iter = iter.next
	}
	iter.value = value
	return nil
}

func (list *LinkedList) Size() int {
	return list.inserted
}
