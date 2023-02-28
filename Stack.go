package dsa

// We need a starting point for our Stack
type Stack struct {
	head *Node
}

// In order to implement our push method for the stack
// we will set our receiver object to be a pointer in
// order for our changes to take effect on the original
// stack instance rather than a copy
func (s *Stack) Push(value int) {
	// first we create an instance of a Node
	// we initialize it's value and make the
	// new Node point to the previous Node in
	// the stack.
	newNode := &Node{value: value, next: s.head}
	s.head = newNode
}

// When you pop from a stack, the value that was popped gets returned
func (s *Stack) Pop() int {
	// first case with popping from a stack that we have to consider
	// is that the stack might be empty in that case we just return 0
	if s.head == nil {
		return 0
	}

	// If the stack is not empty then we need to get the value of
	// the last element that was pushed into the stack.
	value := s.head.value
	s.head = s.head.next
	return value
}
