package main

import (
	"math"
	"testing"
)

func TestUniqueValue(t *testing.T) {
	uniqueArr := []string{"a", "b", "c", "d"}

	if !isUniqueArray(uniqueArr) {
		t.Error("want: true")
	}

	notUniqueArr := []string{"a", "b", "c", "d", "b"}

	if isUniqueArray(notUniqueArr) {
		t.Error("want: false")
	}

}

// Time: O(n), Space: O(n)
func isUniqueArray(arr []string) bool {
	set := newHashSet()

	for _, v := range arr {
		if set.Has(v) {
			return false
		}
		set.Set(v)
	}

	return true
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
