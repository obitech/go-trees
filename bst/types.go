package bst

import (
	"sync"
)

// BSTree represents a binary search tree with a root node and Mutex to protect
// concurrent access.
type BSTree struct {
	sync.RWMutex
	root *node
}

type node struct {
	key     int64
	left    *node
	right   *node
	parent  *node
	payload interface{}
}
