package main

import "testing"

func TestCartFinalize(t *testing.T) {
	cart := NewCart(10)
	cart.Push("Banana")
	cart.Push("Apple")
	cart.Push("Orange")

	if cart.Peek() != "Orange" {
		t.Errorf("Expected: Orange, Got: %s", cart.Peek())
	}

	finalize := cart.Finalize()
	if finalize.count != 3 {
		t.Errorf("Expected: 3, got %d", finalize.count)
	}
	if finalize.Pop() != "Banana" {
		t.Errorf("Expected: Banana, Got: %s", finalize.Pop())
	}
	if finalize.Pop() != "Apple" {
		t.Errorf("Expected: Apple, gGt :%s", finalize.Pop())
	}

	if finalize.Pop() != "Orange" {
		t.Errorf("Expected: Orange, Got: %s", finalize.Pop())
	}

	if finalize.count != 0 {
		t.Errorf("Expected: 0, Got: %d", finalize.count)
	}
}

type Node struct {
	data string
	next *Node
}

func NewNode(data string) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

type Cart struct {
	top     *Node
	count   int
	maxSize int
}

func NewCart(maxSize int) Cart {
	return Cart{
		top:     nil,
		count:   0,
		maxSize: maxSize,
	}
}

func (c *Cart) Push(data string) {
	if c.count >= c.maxSize {
		panic("Stack Overflow")
	}

	newNode := NewNode(data)
	newNode.next = c.top
	c.top = newNode
	c.count++
}

func (c *Cart) Pop() string {
	if c.count <= 0 {
		panic("Stack Underflow")
	}

	data := c.top.data
	c.top = c.top.next
	c.count--
	return data
}

func (c *Cart) Peek() string {
	if c.count <= 0 {
		panic("Stack Underflow")
	}
	return c.top.data
}

func (c *Cart) Finalize() Cart {
	finalize := NewCart(10)

	for c.count > 0 {
		finalize.Push(c.Pop())
	}

	return finalize
}
