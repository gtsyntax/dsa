package dsa

// First we need to create a representation of a node
// which is going to store a value and a reference to
// the next node for our data structure
type Node struct {
	value int
	next  *Node
}
