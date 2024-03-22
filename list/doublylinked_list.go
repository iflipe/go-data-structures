package list

import (
	"errors"
)

type DoublyLinkedList struct {
	head, tail *MegaNode
	inserted   int
}

type MegaNode struct {
	Value int
	Prev  *MegaNode
	Next  *MegaNode
}

func (list *DoublyLinkedList) Add(value int) {
	newNode := &MegaNode{value, nil, nil}
	if list.inserted == 0 {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.Prev = list.tail
		list.tail.Next = newNode
		list.tail = newNode
	}
	list.inserted++
}

func (list *DoublyLinkedList) AddOnIndex(val int, index int) error {
	if index < 0 || index > list.inserted {
		return errors.New("index out of bounds")
	}
	newNode := &MegaNode{val, nil, nil}

	if list.inserted == 0 || index == list.inserted {
		list.Add(val)
		return nil
	} else if index == 0 {
		newNode.Next = list.head
		list.head = newNode
	} else {
		cur := list.head
		if list.inserted/2 > index {
			for i := 0; i < index-1; i++ {
				cur = cur.Next
			}
		} else {
			cur = list.tail
			for i := list.inserted; i > index; i-- {
				cur = cur.Prev
			}
		}
		newNode.Next = cur.Next
		newNode.Prev = cur
		cur.Next = newNode
	}
	list.inserted++
	return nil
}

func (list *DoublyLinkedList) RemoveOnIndex(index int) error {
	if index < 0 || index >= list.inserted {
		return errors.New("index out of bounds")
	}
	if list.inserted == 1 {
		list.head = nil
		list.tail = nil
		list.inserted--
		return nil
	}
	if index == 0 {
		list.head = list.head.Next
		list.head.Prev = nil
	} else if index == list.inserted-1 {
		list.tail = list.tail.Prev
		list.tail.Next = nil
	} else {
		cur := list.head
		if list.inserted/2 > index {
			for i := 0; i < index-1; i++ {
				cur = cur.Next
			}
		} else {
			cur = list.tail
			for i := list.inserted; i > index; i-- {
				cur = cur.Prev
			}
		}
		cur.Next.Next.Prev = cur
		cur.Next = cur.Next.Next
	}
	list.inserted--
	return nil

}

func (list *DoublyLinkedList) Get(index int) (int, error) {
	if index < 0 || index >= list.inserted {
		return index, errors.New("index out of bounds")
	}
	cur := list.head
	if list.inserted/2 > index {
		for i := 0; i < index; i++ {
			cur = cur.Next
		}
	} else {
		cur = list.tail
		for i := list.inserted - 1; i > index; i-- {
			cur = cur.Prev
		}
	}
	return cur.Value, nil
}

func (list *DoublyLinkedList) Set(value int, index int) error {
	if index < 0 || index >= list.inserted {
		return errors.New("index out of bounds")
	}
	cur := list.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Value = value
	return nil
}

func (list *DoublyLinkedList) Size() int {
	return list.inserted
}
