package main

import (
	"slices"
	"testing"
)

func TestInsertAtBeginning(t *testing.T) {
	arr := NewArray()
	arr.Push(2)
	arr.Push(3)
	arr.Push(4)
	arr.Push(5)
	arr.Insert(0, 1)

	want := []int{1, 2, 3, 4, 5}

	if !slices.Equal(arr.elements[:arr.length], want) {
		t.Errorf("got: %v, want: %v", arr.elements, want)
	}
}

const DEFAULT_CAPACITY = 53

func NewArray() *array {
	return &array{
		elements: make([]int, DEFAULT_CAPACITY),
		length:   0,
		capacity: DEFAULT_CAPACITY,
	}
}

// Custom Array
type array struct {
	elements []int
	length   int
	capacity int
}

// Insert a value at a given index
func (c *array) Insert(index int, value int) {
	if index < 0 || index > c.length {
		panic("Index out of bounds")
	}

	// Grow if capacity is full
	if c.length == c.capacity {
		c.grow()
	}

	// O(1)
	if index == c.length {
		c.elements[index] = value
		c.length++
		return
	}

	// O(n)
	for i := c.length; i > index; i-- {
		c.elements[i] = c.elements[i-1]
	}
	c.elements[index] = value
	c.length++
}

// Push a value at the end of the array
func (c *array) Push(value int) {
	c.elements[c.length] = value
	c.length++
}

// Resize the array
func (c *array) resize(newCapacity int) {
	if newCapacity == c.capacity {
		return
	}

	newElements := make([]int, newCapacity)
	//Space: O(n)
	for i := 0; i < c.capacity; i++ {
		newElements[i] = c.elements[i]
	}

	c.elements = newElements
	c.capacity = newCapacity
}

// Increase the size of the array
func (c *array) grow() {
	newCapacity := 2 * c.capacity
	c.resize(newCapacity)
}
