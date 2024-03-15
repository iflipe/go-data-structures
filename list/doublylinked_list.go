package list

import (
	"errors"
)

type DoublyLinkedList struct {
	head, tail *MegaNode
	inserted   int
}

type MegaNode struct {
	value int
	prev  *MegaNode
	next  *MegaNode
}

func (list *DoublyLinkedList) Add(value int) {
	newNode := &MegaNode{value, nil, nil}
	if list.inserted == 0 {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
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
		newNode.next = list.head
		list.head = newNode
	} else {
		cur := list.head
		if list.inserted/2 > index {
			for i := 0; i < index-1; i++ {
				cur = cur.next
			}
		} else {
			cur = list.tail
			for i := list.inserted; i > index; i-- {
				cur = cur.prev
			}
		}
		newNode.next = cur.next
		newNode.prev = cur
		cur.next = newNode
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
		list.head = list.head.next
		list.head.prev = nil
	} else if index == list.inserted-1 {
		list.tail = list.tail.prev
		list.tail.next = nil
	} else {
		cur := list.head
		if list.inserted/2 > index {
			for i := 0; i < index-1; i++ {
				cur = cur.next
			}
		} else {
			cur = list.tail
			for i := list.inserted; i > index; i-- {
				cur = cur.prev
			}
		}
		cur.next.next.prev = cur
		cur.next = cur.next.next
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
			cur = cur.next
		}
	} else {
		cur = list.tail
		for i := list.inserted - 1; i > index; i-- {
			cur = cur.prev
		}
	}
	return cur.value, nil
}

func (list *DoublyLinkedList) Set(value int, index int) error {
	if index < 0 || index >= list.inserted {
		return errors.New("index out of bounds")
	}
	cur := list.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.value = value
	return nil
}

func (list *DoublyLinkedList) Size() int {
	return list.inserted
}
