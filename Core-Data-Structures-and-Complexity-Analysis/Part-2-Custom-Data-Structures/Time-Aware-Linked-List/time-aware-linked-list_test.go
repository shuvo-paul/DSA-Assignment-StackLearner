package main

import (
	"slices"
	"testing"
	"time"
)

func TestTimeAwareLinkedList(t *testing.T) {
	ta := NewTimeAwareLinkedList()
	ta.Insert("A")
	time.Sleep(5 * time.Millisecond)

	ta.Insert("B")
	time.Sleep(5 * time.Millisecond)

	ta.Insert("C")
	time.Sleep(5 * time.Millisecond)

	ta.Insert("D")

	nodes := ta.RetriveRecent(2 * time.Millisecond)
	if len(nodes) != 1 {
		t.Errorf("Expected: 1, Got: %d", len(nodes))
	}

	if nodes[0].value != "D" {
		t.Errorf("Expected: D, Got: %s", nodes[0].value)
	}
	nodes = ta.RetriveRecent(6 * time.Millisecond)
	if len(nodes) != 2 {
		t.Errorf("Expected: 2, Got: %d", len(nodes))
	}
	if !slices.Equal([]string{"D", "C"}, []string{nodes[0].value, nodes[1].value}) {
		t.Errorf("Expected: [D, C], Got: %s", []string{nodes[0].value, nodes[1].value})
	}

	nodes = ta.RetriveRecent(14 * time.Millisecond)
	if len(nodes) != 3 {
		t.Errorf("Expected: 3, Got: %d", len(nodes))
	}
	if !slices.Equal([]string{"D", "C", "B"}, []string{nodes[0].value, nodes[1].value, nodes[2].value}) {
		t.Errorf("Expected: [D, C, B], Got: %s", []string{nodes[0].value, nodes[1].value, nodes[2].value})
	}
}

func NewNode(value string) *Node {
	return &Node{
		value: value,
		time:  time.Now(),
		next:  nil,
	}
}

type Node struct {
	value string
	time  time.Time
	next  *Node
}

func NewTimeAwareLinkedList() *TimeAwareLinkedList {
	return &TimeAwareLinkedList{
		head:  nil,
		count: 0,
	}
}

type TimeAwareLinkedList struct {
	head  *Node
	count int
}

func (ta *TimeAwareLinkedList) Insert(value string) {
	node := NewNode(value)
	ta.count++
	node.next = ta.head
	ta.head = node
}

func (ta *TimeAwareLinkedList) Remove(value string) {
	current := ta.head
	prev := ta.head

	for current != nil {
		if current.value == value {
			if current == ta.head {
				ta.head = ta.head.next
			}

			if prev != nil {
				prev.next = current.next
			}
			ta.count--
			return
		}

		prev = current
		current = current.next
	}
}

func (ta *TimeAwareLinkedList) RetriveRecent(duration time.Duration) []*Node {
	result := make([]*Node, 0)
	currentTime := time.Now()
	current := ta.head

	for current != nil {
		if currentTime.Sub(current.time) <= duration {
			result = append(result, current)
		}
		current = current.next
	}
	return result
}
