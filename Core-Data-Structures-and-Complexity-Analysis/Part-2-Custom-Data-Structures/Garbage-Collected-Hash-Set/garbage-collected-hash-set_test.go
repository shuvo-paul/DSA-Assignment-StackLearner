package main

import (
	"math"
	"sync"
	"testing"
	"time"
)

func TestHashSet(t *testing.T) {
	hs := NewHashSet(50 * time.Millisecond)
	go hs.CleanUpWorker()
	defer hs.StopCleanupWorker()

	hs.Insert("A")
	if !hs.Find("A") {
		t.Error("Expected: True, Got: False")
	}

	time.Sleep(55 * time.Millisecond)
	if hs.Find("A") {
		t.Error("Expected: False, Got: true")
	}
}

func newNode(value string, ttl time.Duration) *node {
	return &node{
		value:      value,
		next:       nil,
		expiration: time.Now().Add(ttl),
	}
}

type node struct {
	value      string
	next       *node
	expiration time.Time
}

func NewBucket() *bucket {
	return &bucket{
		head:  nil,
		count: 0,
	}
}

type bucket struct {
	head  *node
	count int
}

func (b *bucket) insert(node *node) {
	b.count++
	if b.head == nil {
		b.head = node
		return
	}

	node.next = b.head
	b.head = node
}

func (b *bucket) remove(node *node) {
	current := b.head
	prev := b.head
	for current != nil {
		if current != node {
			prev = current
			current = current.next
			continue
		}

		b.count--

		if b.head == current {
			b.head = current.next
			return
		}

		if prev != nil {
			prev.next = current.next
		}
	}
}

func (b *bucket) find(value string) *node {
	current := b.head
	for current != nil {
		if current.value == value {
			return current
		}
		current = current.next
	}
	return nil
}

func (b *bucket) entries() []node {
	nodes := []node{}

	current := b.head

	for current != nil {
		nodes = append(nodes, *current)
	}

	return nodes
}

const DEFAULT_SIZE = 53

type HashSet struct {
	mu       sync.RWMutex
	capacity int
	table    []*bucket
	count    int
	ttl      time.Duration
	stopChan chan struct{}
}

func NewHashSet(ttl time.Duration) HashSet {
	return HashSet{
		capacity: DEFAULT_SIZE,
		table:    make([]*bucket, DEFAULT_SIZE),
		count:    0,
		ttl:      ttl,
		stopChan: make(chan struct{}),
	}
}

func (hs *HashSet) hash(key string) int {
	hash := 5381

	for _, v := range key {
		hash = hash ^ int(v)
	}
	return int(math.Abs(float64(hash))) % hs.capacity
}

func (hs *HashSet) Insert(value string) {
	hs.mu.Lock()
	defer hs.mu.Unlock()
	index := hs.hash(value)
	if hs.table[index] == nil {
		hs.table[index] = NewBucket()
	}

	if node := hs.table[index].find(value); node != nil {
		node.expiration = time.Now().Add(hs.ttl)
		return
	}

	newNode := newNode(value, hs.ttl)

	hs.table[index].insert(newNode)
	hs.count++
}

// O(1)
func (hs *HashSet) Remove(value string) {
	hs.mu.Lock()
	defer hs.mu.Unlock()

	index := hs.hash(value)
	if hs.table[index] == nil {
		return
	}

	node := hs.table[index].find(value)

	if node == nil {
		return
	}

	hs.table[index].remove(node)
	hs.count--
}

// O(1)
func (hs *HashSet) Find(value string) bool {
	hs.mu.RLock()
	defer hs.mu.RUnlock()

	index := hs.hash(value)
	if hs.table[index] == nil {
		return false
	}

	node := hs.table[index].find(value)

	if node == nil {
		return false
	}

	if time.Now().After(node.expiration) {
		return false
	}

	node.expiration = time.Now().Add(hs.ttl)
	return true
}

func (hs *HashSet) Entries() []node {
	hs.mu.RLock()
	defer hs.mu.RUnlock()
	nodes := []node{}
	for _, bucket := range hs.table {
		if bucket == nil {
			continue
		}
		nodes = append(nodes, bucket.entries()...)
	}
	return nodes
}

func (hs *HashSet) Cleanup() {
	nodes := hs.Entries()

	for _, node := range nodes {
		if time.Now().After(node.expiration) {
			hs.Remove(node.value)
		}
	}
}

func (hs *HashSet) CleanUpWorker() {
	ticker := time.NewTicker(hs.ttl / 10)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			hs.Cleanup()
		case <-hs.stopChan:
			return
		}
	}
}

func (hs *HashSet) StopCleanupWorker() {
	close(hs.stopChan)
}
