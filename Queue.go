package dsa

type Queue struct {
	head *Node
	tail *Node
}

func (q *Queue) Enqueue(value int) {
	// firstly we construct a node that will be enqueued
	newNode := &Node{value: value}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

func (q *Queue) Dequeue() int {
	if q.head == nil {
		return 0
	}
	value := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return value
}
