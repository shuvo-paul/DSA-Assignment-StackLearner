package main

import (
	"math"
	"testing"
)

func TestRehash(t *testing.T) {
	hashTable := newHashTable()

	hashTable.Set("name", "Jhon Doe")
	hashTable.Set("city", "khulna")
	hashTable.Set("country", "Bangladesh")
	hashTable.Set("age", "35")

	hashTable.resize(5)

	if cap(hashTable.table) != 5 {
		t.Errorf("Want: 5, Got: %d", cap(hashTable.table))
	}

	hashTable.Set("height", "5.9 Inch")
	hashTable.Set("weight", "70KG")

	if cap(hashTable.table) <= 5 {
		t.Error("Capacity should be learger now!")
	}
}

func newNode(key, value string) *node {
	return &node{
		key:   key,
		value: value,
		next:  nil,
	}
}

type node struct {
	key   string
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

func (b *bucket) insert(key, value string) {
	newNode := newNode(key, value)
	if b.head == nil {
		b.head = newNode
		return
	}

	newNode.next = b.head
	b.head = newNode
}

func (b *bucket) find(key string) string {
	current := b.head
	for current != nil {
		if current.key == key {
			return current.value
		}

		current = current.next
	}

	return ""
}

// Time: O(n), Space: O(1)
func (b *bucket) remove(key string) {
	if b.head == nil {
		return
	}

	var prevNode *node
	current := b.head

	for current != nil {
		if current.key == key {
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

func newHashTable() *hashTable {
	return &hashTable{
		count: 0,
		table: make([]*bucket, DEFAULT_CAPACITY),
	}
}

type hashTable struct {
	count int
	table []*bucket
}

func (h *hashTable) Set(key, value string) {
	length := cap(h.table)
	if float64(h.count/length) >= 0.5 {
		h.resize(length * 2)
	}

	if float64(h.count/length) <= 0.4 && length > DEFAULT_CAPACITY {
		h.resize(length / 2)
	}
	index := h.hash(key)

	if h.table[index] == nil {
		h.table[index] = newBucket()
	}

	bucket := h.table[index]

	bucket.insert(key, value)
	h.count++
}

func (h *hashTable) Get(key string) string {
	index := h.hash(key)

	return h.table[index].find(key)
}

func (h *hashTable) Delete(key string) {
	index := h.hash(key)
	if h.table[index] == nil {
		return
	}
	h.table[index].remove(key)
	h.count--
}

// Time: O(n) , Space: O(n)
func (h *hashTable) resize(cap int) {
	entries := h.entries()
	h.table = make([]*bucket, cap)
	h.count = 0
	for _, node := range entries {
		h.Set(node.key, node.value)
	}
}

func (h *hashTable) entries() []*node {
	entries := make([]*node, 0)
	for _, bucket := range h.table {
		if bucket == nil {
			continue
		}

		entries = append(entries, bucket.entries()...)
	}

	return entries
}

func (h *hashTable) hash(key string) int {
	hash := 5381

	for _, v := range key {
		hash = hash ^ int(v)
	}

	return int(math.Abs(float64(hash))) % cap(h.table)
}
