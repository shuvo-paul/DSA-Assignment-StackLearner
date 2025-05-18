package main

import (
	"math"
	"testing"
)

func TestSearchPatient(t *testing.T) {
	hm := NewHashMap()
	hm.Set("1", "Alice")
	hm.Set("2", "Bob")
	hm.Set("3", "Charlie")
	patient, ok := hm.Get("2")
	if !ok || patient.name != "Bob" {
		t.Errorf("Expected Bob, got %s", patient.name)
	}
}

func newPatient(id string, name string) *patient {
	return &patient{
		id:   id,
		name: name,
		next: nil,
	}
}

type patient struct {
	id       string
	name     string
	quantity int
	price    float64
	next     *patient
}

func newBucket() *bucket {
	return &bucket{
		head: nil,
	}
}

type bucket struct {
	head *patient
}

func (b *bucket) insert(id string, name string) bool {
	for patient := b.head; patient != nil; patient = patient.next {
		if patient.id == id {
			patient.name = name
			return false
		}
	}

	patient := newPatient(id, name)
	if b.head == nil {
		b.head = patient
		return true
	}

	patient.next = b.head
	b.head = patient
	return true
}

func (b *bucket) find(id string) (patient, bool) {
	current := b.head
	for current != nil {
		if current.id == id {
			return *current, true
		}

		current = current.next
	}

	return patient{}, false
}

// Time: O(n), Space: O(1)
func (b *bucket) remove(id string) bool {
	if b.head == nil {
		return false
	}

	var prevNode *patient
	current := b.head

	for current != nil {
		if current.id == id {
			if prevNode == nil {
				b.head = current.next
			} else {
				prevNode.next = current.next
			}

			return true
		}

		prevNode = current
		current = current.next
	}

	return false
}

// Time: O(n), Space: O(1)
func (b *bucket) entries() []patient {
	current := b.head
	entries := make([]patient, 0)
	for current != nil {
		entries = append(entries, *current)
		current = current.next
	}

	return entries
}

const DEFAULT_CAPACITY = 53

func NewHashMap() *hashMap {
	return &hashMap{
		count: 0,
		table: make([]*bucket, DEFAULT_CAPACITY),
	}
}

type hashMap struct {
	count int
	table []*bucket
}

func (hm *hashMap) Set(id string, name string) {
	length := cap(hm.table)
	if float64(hm.count)/float64(length) >= 0.7 {
		hm.resize(length * 2)
	} else if float64(hm.count)/float64(length) <= 0.3 && length > DEFAULT_CAPACITY {
		hm.resize(length / 2)
	}
	index := hm.hash(id)

	if hm.table[index] == nil {
		hm.table[index] = newBucket()
	}

	bucket := hm.table[index]

	if bucket.insert(id, name) {
		hm.count++
	}
}

func (hm *hashMap) Get(key string) (patient, bool) {
	index := hm.hash(key)
	if hm.table[index] == nil {
		return patient{}, false
	}
	return hm.table[index].find(key)
}

func (hm *hashMap) Delete(key string) {
	index := hm.hash(key)
	if hm.table[index] == nil {
		return
	}
	if hm.table[index].remove(key) {
		hm.count--
	}
}

// Time: O(n) , Space: O(n)
func (hm *hashMap) resize(cap int) {
	entries := hm.entries()
	hm.table = make([]*bucket, cap)
	hm.count = 0
	for _, patient := range entries {
		index := hm.hash(patient.id)
		if hm.table[index] == nil {
			hm.table[index] = newBucket()
		}
		hm.table[index].insert(patient.id, patient.name)
		hm.count++
	}
}

func (hm *hashMap) entries() []patient {
	entries := make([]patient, 0)
	for _, bucket := range hm.table {
		if bucket == nil {
			continue
		}

		entries = append(entries, bucket.entries()...)
	}

	return entries
}

func (hm *hashMap) hash(key string) int {
	hash := 5381

	for _, v := range key {
		hash = hash ^ int(v)
	}

	return int(math.Abs(float64(hash))) % cap(hm.table)
}
