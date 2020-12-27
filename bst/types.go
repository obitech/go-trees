package bst

import (
	"sync"
)

// Tree represents a binary search tree with a root not and Mutex to protect
// concurrent access.
type Tree struct {
	sync.RWMutex
	root *node
}

type Item struct {
	Key     int64
	Payload interface{}
}

type node struct {
	key     int64
	left    *node
	right   *node
	parent  *node
	payload interface{}
}
