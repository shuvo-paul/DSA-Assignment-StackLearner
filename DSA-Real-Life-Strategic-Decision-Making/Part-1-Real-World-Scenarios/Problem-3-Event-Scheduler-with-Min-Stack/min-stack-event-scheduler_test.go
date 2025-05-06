package main

import "testing"

func TestMinStack(t *testing.T) {
	ms := NewMinStack()
	ms.Push(6)
	ms.Push(8)
	ms.Push(7)
	ms.Push(5)
	ms.Push(10)

	if ms.GetMin() != 5 {
		t.Errorf("want: 5, Got: %d", ms.GetMin())
	}

	ms.Pop()
	ms.Pop()

	if ms.GetMin() != 6 {
		t.Errorf("want: 6, Got: %d", ms.GetMin())
	}

	ms.Push(9)
	ms.Push(4)
	ms.Push(7)

	if ms.GetMin() != 4 {
		t.Errorf("want: 4, Got: %d", ms.GetMin())
	}
}

type Node struct {
	data int
	min  int
	next *Node
}

func NewNode(data int) *Node {
	return &Node{
		data: data,
		min:  data,
		next: nil,
	}
}

type MinStack struct {
	top *Node
}

func NewMinStack() MinStack {
	return MinStack{
		top: nil,
	}
}

func (ms *MinStack) Push(x int) {
	newNode := NewNode(x)
	if ms.top == nil {
		ms.top = newNode
		return
	}

	if x > ms.top.min {
		newNode.min = ms.top.min
	}

	newNode.next = ms.top
	ms.top = newNode
}

func (ms *MinStack) Pop() int {
	if ms.top == nil {
		panic("Stack is empty")
	}

	data := ms.top.data
	ms.top = ms.top.next
	return data
}

func (ms *MinStack) GetMin() int {
	if ms.top == nil {
		panic("Stack is empty")
	}

	return ms.top.min
}
