package main

import "testing"

func TestMatchMaking(t *testing.T) {
	queue := NewQueue(5)
	queue.Enqueue("Player 1")
	queue.Enqueue("Player 2")
	queue.Enqueue("Player 3")
	queue.Enqueue("Player 4")
	queue.Enqueue("Player 5")
	players := queue.MakeMatch(2)
	if len(players) != 2 {
		t.Errorf("Expected: 2, Got: %d", len(players))
	}

	if players[0] != "Player 1" || players[1] != "Player 2" {
		t.Errorf("Expected: Player 1, Got: %s", players[0])
		t.Errorf("Expected: Player 2, Got: %s", players[1])
	}

	queue.Enqueue("Player 6")
	queue.Enqueue("Player 7")
	players = queue.MakeMatch(2)
	if len(players) != 2 {
		t.Errorf("Expected: 2, Got: %d", len(players))
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

type Queue struct {
	front *Node
	rear  *Node
}

func NewQueue(capacity int) Queue {
	return Queue{
		front: nil,
		rear:  nil,
	}
}

func (q *Queue) Enqueue(data string) {
	newNode := NewNode(data)

	if q.front == nil {
		q.front = newNode
		q.rear = newNode
		return
	}

	q.rear.next = newNode
	q.rear = newNode
}

func (q *Queue) Dequeue() string {
	if q.front == nil {
		panic("Queue is empty")
	}

	value := q.front.data
	q.front = q.front.next

	return value
}

// Time: O(n); Space: O(n)
func (q *Queue) MakeMatch(numOfEl int) []string {
	players := make([]string, numOfEl)

	for i := range numOfEl {
		players[i] = q.Dequeue()
	}

	return players
}
