package main

import "testing"

func TestDeleteByValue(t *testing.T) {
	linkedList := newLinkedList()
	linkedList.Append(9)
	linkedList.Append(8)
	linkedList.Append(7)
	linkedList.Append(6)

	linkedList.Remove(8)

	if linkedList.size != 3 {
		t.Error("Failed to delete node")
	}

	linkedList.Remove(6)

	if linkedList.tail.data == 6 {
		t.Error("Failed to delete tail node")
	}

	linkedList.Remove(9)
	if linkedList.head.data == 9 {
		t.Error("Failed to delete head node")
	}

	linkedList.Remove(7)

	if linkedList.head != nil {
		t.Error("Head should be nil")
	}
}

func newNode(data int) *node {
	return &node{
		data: data,
		next: nil,
	}
}

type node struct {
	data int
	next *node
}

func newLinkedList() *linkedList {
	return &linkedList{
		head: nil,
		size: 0,
		tail: nil,
	}
}

type linkedList struct {
	head *node
	size int
	tail *node
}

func (l *linkedList) Append(data int) {
	newNode := newNode(data)
	l.size++

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	l.tail.next = newNode
	l.tail = newNode

}

func (l *linkedList) Prepend(data int) {
	newNode := newNode(data)
	newNode.next = l.head
	l.head = newNode
	l.size++
}

func (l *linkedList) Remove(data int) {
	if l.head == nil {
		panic("Empty head")
	}

	var prevNode *node
	current := l.head

	// Time: O(n), Space: O(1)
	for range l.size {

		if current.data != data {
			prevNode = current
			current = current.next
			continue
		}

		if prevNode != nil {
			prevNode.next = current.next
		}

		if current == l.head {
			l.head = l.head.next
		}

		if current == l.tail {
			l.tail = prevNode
		}
		l.size--
		break

	}
}
