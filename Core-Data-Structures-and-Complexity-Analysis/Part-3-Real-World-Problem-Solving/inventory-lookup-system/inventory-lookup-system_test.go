package main

import (
	"fmt"
	"math"
	"testing"
)

func TestInventory_AddAndGet(t *testing.T) {
	inv := NewInventory()

	inv.Set("sku123", "Laptop", 5, 1299.99)
	inv.Set("sku124", "Mouse", 20, 25.50)

	p1, ok1 := inv.Get("sku123")
	if !ok1 || p1.name != "Laptop" || p1.quantity != 5 || p1.price != 1299.99 {
		t.Errorf("Failed to retrieve sku123: got %+v, ok=%v", p1, ok1)
	}

	p2, ok2 := inv.Get("sku124")
	if !ok2 || p2.name != "Mouse" || p2.quantity != 20 || p2.price != 25.50 {
		t.Errorf("Failed to retrieve sku124: got %+v, ok=%v", p2, ok2)
	}
}

func TestInventory_Update(t *testing.T) {
	inv := NewInventory()

	inv.Set("sku200", "Keyboard", 10, 45.0)
	inv.Set("sku200", "Keyboard", 15, 42.0) // update same ID

	p, ok := inv.Get("sku200")
	if !ok || p.quantity != 15 || p.price != 42.0 {
		t.Errorf("Update failed for sku200: got %+v", p)
	}
}

func TestInventory_Delete(t *testing.T) {
	inv := NewInventory()

	inv.Set("sku300", "Monitor", 7, 250.0)
	inv.Delete("sku300")

	_, ok := inv.Get("sku300")
	if ok {
		t.Errorf("Expected sku300 to be deleted")
	}

	inv.Delete("nonexistent")
}

func TestInventory_Resize(t *testing.T) {
	inv := NewInventory()

	for i := range 70 {
		id := fmt.Sprintf("sku%d", i)
		inv.Set(id, "Item", i+1, float64(i*10))
	}

	for i := range 70 {
		id := fmt.Sprintf("sku%d", i)
		_, ok := inv.Get(id)
		if !ok {
			t.Errorf("Missing entry after resize: %s", id)
		}
	}
}

func TestInventory_GetNonExistent(t *testing.T) {
	inv := NewInventory()

	_, ok := inv.Get("notfound")
	if ok {
		t.Errorf("Expected notfound to be absent")
	}
}

func newProduct(id string, name string, quantity int, price float64) *product {
	return &product{
		id:       id,
		name:     name,
		quantity: quantity,
		price:    price,
		next:     nil,
	}
}

type product struct {
	id       string
	name     string
	quantity int
	price    float64
	next     *product
}

func newBucket() *bucket {
	return &bucket{
		head: nil,
	}
}

type bucket struct {
	head *product
}

func (b *bucket) insert(id string, name string, quantity int, price float64) bool {
	for product := b.head; product != nil; product = product.next {
		if product.id == id {
			product.name = name
			product.quantity = quantity
			product.price = price
			return false
		}
	}

	product := newProduct(id, name, quantity, price)
	if b.head == nil {
		b.head = product
		return true
	}

	product.next = b.head
	b.head = product
	return true
}

func (b *bucket) find(id string) (product, bool) {
	current := b.head
	for current != nil {
		if current.id == id {
			return *current, true
		}

		current = current.next
	}

	return product{}, false
}

// Time: O(n), Space: O(1)
func (b *bucket) remove(id string) bool {
	if b.head == nil {
		return false
	}

	var prevNode *product
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
func (b *bucket) entries() []product {
	current := b.head
	entries := make([]product, 0)
	for current != nil {
		entries = append(entries, *current)
		current = current.next
	}

	return entries
}

const DEFAULT_CAPACITY = 53

func NewInventory() *inventory {
	return &inventory{
		count: 0,
		table: make([]*bucket, DEFAULT_CAPACITY),
	}
}

type inventory struct {
	count int
	table []*bucket
}

func (inv *inventory) Set(id string, name string, quantity int, price float64) {
	length := cap(inv.table)
	if float64(inv.count)/float64(length) >= 0.7 {
		inv.resize(length * 2)
	} else if float64(inv.count)/float64(length) <= 0.3 && length > DEFAULT_CAPACITY {
		inv.resize(length / 2)
	}
	index := inv.hash(id)

	if inv.table[index] == nil {
		inv.table[index] = newBucket()
	}

	bucket := inv.table[index]

	if bucket.insert(id, name, quantity, price) {
		inv.count++
	}
}

func (inv *inventory) Get(key string) (product, bool) {
	index := inv.hash(key)
	if inv.table[index] == nil {
		return product{}, false
	}
	return inv.table[index].find(key)
}

func (inv *inventory) Delete(key string) {
	index := inv.hash(key)
	if inv.table[index] == nil {
		return
	}
	if inv.table[index].remove(key) {
		inv.count--
	}
}

// Time: O(n) , Space: O(n)
func (inv *inventory) resize(cap int) {
	entries := inv.entries()
	inv.table = make([]*bucket, cap)
	inv.count = 0
	for _, product := range entries {
		index := inv.hash(product.id)
		if inv.table[index] == nil {
			inv.table[index] = newBucket()
		}
		inv.table[index].insert(product.id, product.name, product.quantity, product.price)
		inv.count++
	}
}

func (inv *inventory) entries() []product {
	entries := make([]product, 0)
	for _, bucket := range inv.table {
		if bucket == nil {
			continue
		}

		entries = append(entries, bucket.entries()...)
	}

	return entries
}

func (inv *inventory) hash(key string) int {
	hash := 5381

	for _, v := range key {
		hash = hash ^ int(v)
	}

	return int(math.Abs(float64(hash))) % cap(inv.table)
}
