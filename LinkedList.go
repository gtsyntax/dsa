package main

// First we need to create a representation of a node
// which is going to store a value and a reference to
// the next node in the linkedlist.
type Node struct {
	value int
	next  *Node
}

// Next we need a starting point for our linkedlist construction
type List struct {
	head *Node
}

// The Add method will be used to add elements to our linkedlist object
// we make the receiver a pointer because we want to apply the method
// on the original object and not a copy
func (l *List) Add(value int) {
	// first we create a new Node object
	newNode := &Node{value: value}

	// next we have to check if the head is nil or not
	// if the head is nil that means newNode will be the
	// first Node added to List however if head isn't nil
	// then we have to traverse to the end of the List and
	// add newNode at the end.
	if l.head == nil {
		l.head = newNode
	} else {
		curr := l.head
		for curr.next != nil {
			// if the next node is not nil then we set it
			// to be our current position
			curr = curr.next
		}
		// once we are at the last element in the List we
		// set the current node's next pointer to point to our newNode
		curr.next = newNode
	}
}
