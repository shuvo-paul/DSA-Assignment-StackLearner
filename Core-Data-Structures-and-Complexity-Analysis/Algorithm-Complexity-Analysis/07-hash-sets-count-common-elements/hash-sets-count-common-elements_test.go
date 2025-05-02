package main

import (
	"math"
	"testing"
)

func TestCommonElements(t *testing.T) {
	setA := newHashSet()
	setA.Set("apple")
	setA.Set("banana")
	setA.Set("cherry")

	setB := newHashSet()
	setB.Set("banana")
	setB.Set("orange")
	setB.Set("grape")
	commonCount := countCommonElements(setA, setB)
	if commonCount != 1 {
		t.Errorf("Want: 1, Got: %d", commonCount)
	}

	setA.Set("orange")

	commonCount = countCommonElements(setA, setB)
	if commonCount != 2 {
		t.Errorf("Want: 2, Got: %d", commonCount)
	}
}

func countCommonElements(A, B *hashSet) int {
	count := 0
	for _, v := range A.Values() {
		if B.Has(v) {
			count++
		}
	}

	return count
}

func newNode(value string) *node {
	return &node{
		value: value,
		next:  nil,
	}
}

type node struct {
	value string
	next  *node
}

func newBucket() *bucket {
	return &bucket{
		head: nil,
		size: 0,
	}
}

type bucket struct {
	head *node
	size int
}

func (b *bucket) Add(value string) {
	newNode := newNode(value)

	if b.head != nil {
		newNode.next = b.head
	}

	b.head = newNode
	b.size++
}

func (b *bucket) Has(value string) bool {
	current := b.head

	//Time: O(n), Space: O(1)
	for {
		if current == nil {
			break
		}

		if current.value == value {
			return true
		}

		current = current.next
	}

	return false
}

func (b *bucket) Values() []string {
	current := b.head
	elements := make([]string, 0)
	for current != nil {
		elements = append(elements, current.value)
		current = current.next
	}
	return elements
}

func newHashSet() *hashSet {
	return &hashSet{
		table: make([]*bucket, 10),
	}
}

type hashSet struct {
	table []*bucket
}

func (h *hashSet) hash(value string) int {
	hash := 5381

	for _, ch := range value {
		hash = (hash * 33) ^ int(ch)
	}

	return int(math.Abs(float64(hash))) % 10
}

// Time: O(1), Space: O(1)
func (h *hashSet) Set(value string) {
	index := h.hash(value)

	if h.table[index] == nil {
		h.table[index] = newBucket()
	}

	bucket := h.table[index]

	bucket.Add(value)
}

// Time: O(1), Space: O(1)
func (h *hashSet) Has(value string) bool {
	index := h.hash(value)
	if h.table[index] == nil {
		return false
	}

	return h.table[index].Has(value)
}

func (h *hashSet) Values() []string {
	values := make([]string, 0)

	// Time: O(n), Space: O(1)
	for _, bucket := range h.table {
		if bucket != nil {
			for _, v := range bucket.Values() {
				values = append(values, v)
			}
		}
	}

	return values
}
