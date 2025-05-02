package main

import (
	"math"
	"testing"
)

func BenchmarkArrayLookup(b *testing.B) {
	arr := []string{"banana", "papaya", "gelato", "poopaye"}
	target := "poopaye"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range arr {
			if v == target {
				break
			}
		}
	}
}

func BenchmarkHashLookup(b *testing.B) {
	hashSet := newHashSet()
	hashSet.Set("banana")
	hashSet.Set("papaya")
	hashSet.Set("gelato")
	hashSet.Set("poopaye")
	target := "poopaye"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hashSet.Find(target)
	}
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
	}
}

type bucket struct {
	head *node
}

func (b *bucket) insert(value string) {
	newNode := newNode(value)
	if b.head == nil {
		b.head = newNode
		return
	}

	newNode.next = b.head
	b.head = newNode
}

func (b *bucket) find(value string) bool {
	current := b.head
	for current != nil {
		if current.value == value {
			return true
		}

		current = current.next
	}

	return false
}

// Time: O(n), Space: O(1)
func (b *bucket) remove(value string) {
	if b.head == nil {
		return
	}

	var prevNode *node
	current := b.head

	for current != nil {
		if current.value == value {
			if prevNode == nil {
				b.head = current.next
			} else {
				prevNode.next = current.next
			}
		}

		prevNode = current
		current = current.next
	}
}

// Time: O(n), Space: O(1)
func (b *bucket) entries() []*node {
	current := b.head
	entries := make([]*node, 0)
	for {
		if current == nil {
			break
		}
		entries = append(entries, current)
		current = current.next
	}

	return entries
}

const DEFAULT_CAPACITY = 11

func newHashSet() *hashSet {
	return &hashSet{
		count: 0,
		table: make([]*bucket, DEFAULT_CAPACITY),
	}
}

type hashSet struct {
	count int
	table []*bucket
}

func (h *hashSet) Set(value string) {
	length := cap(h.table)
	if float64(h.count/length) >= 0.5 {
		h.resize(length * 2)
	}

	if float64(h.count/length) <= 0.4 && length > DEFAULT_CAPACITY {
		h.resize(length / 2)
	}
	index := h.hash(value)

	if h.table[index] == nil {
		h.table[index] = newBucket()
	}

	bucket := h.table[index]

	bucket.insert(value)
	h.count++
}

// Time: O(1), Space: O(1)
func (h *hashSet) Find(key string) bool {
	index := h.hash(key)

	if h.table[index] == nil {
		return false
	}
	return h.table[index].find(key)
}

func (h *hashSet) Delete(value string) {
	index := h.hash(value)
	if h.table[index] == nil {
		return
	}
	h.table[index].remove(value)
	h.count--
}

// Time: O(n) , Space: O(n)
func (h *hashSet) resize(cap int) {
	entries := h.entries()
	h.table = make([]*bucket, cap)
	h.count = 0
	for _, node := range entries {
		h.Set(node.value)
	}
}

// Returns all the entries in the hash table in a slice.
// Time: O(n), Space: O(n)
func (h *hashSet) entries() []*node {
	entries := make([]*node, 0)
	for _, bucket := range h.table {
		if bucket == nil {
			continue
		}

		entries = append(entries, bucket.entries()...)
	}

	return entries
}

func (h *hashSet) hash(key string) int {
	hash := 5381

	for _, v := range key {
		hash = hash ^ int(v)
	}

	return int(math.Abs(float64(hash))) % cap(h.table)
}
