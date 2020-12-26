package bst

// Indexer needs to be implemented by types that are part of the payload of
// the tree. The Index function returns the index, or key, of the node.
type Indexer interface {
	Index() int64
}

type Tree struct {
	Root *Node
}

type Node struct {
	Key     int64
	Left    *Node
	Right   *Node
	Payload Indexer
}
