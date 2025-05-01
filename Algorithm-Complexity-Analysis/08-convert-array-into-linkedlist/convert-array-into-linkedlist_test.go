package main

import (
	"reflect"
	"testing"
)

func TestArrayIntoLinkedList(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	linkedList := convertArrayToLinkedList(arr)
	if linkedList.count != len(arr) {
		t.Errorf("Expected %d, got %d", len(arr), linkedList.count)
	}

	if reflect.TypeOf(linkedList) == reflect.TypeOf(arr) {
		t.Errorf("Expected %T, got %T", linkedList, arr)
	}
}

func convertArrayToLinkedList(arr []int) LinkedList {
	linkedList := newLinkedList()
	for _, v := range arr {
		linkedList.Append(v)
	}

	return linkedList
}

func newNode(data int) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

type Node struct {
	data int
	next *Node
}

func newLinkedList() LinkedList {
	return LinkedList{
		count: 0,
		head:  nil,
		tail:  nil,
	}
}

type LinkedList struct {
	count int
	head  *Node
	tail  *Node
}

func (l *LinkedList) Append(data int) {
	newNode := newNode(data)
	if l.head == nil {
		l.head = newNode
	}

	if l.tail != nil {
		l.tail.next = newNode
	}
	l.tail = newNode
	l.count++
}
