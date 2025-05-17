package main

import "testing"

func TestInsertAtTheEnd(t *testing.T) {
	linkedList := newLinkedList()
	linkedList.Append(9)
	linkedList.Append(8)
	linkedList.Append(7)
	linkedList.Append(6)

	if linkedList.tail.data != 6 {
		t.Error("data mismatched")
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
