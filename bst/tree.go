// Package bst implements a binary search tree with arbitrary payloads.
package bst

import (
	"math"
	"sync"
)

// New returns an empty binary search tree.
func New() *Tree {
	return &Tree{
		RWMutex: sync.RWMutex{},
		root:    nil,
	}
}

// Root returns the payload of the root node of the tree.
func (t *Tree) Root() interface{} {
	t.RLock()
	defer t.RUnlock()

	if t.root == nil {
		return nil
	}

	return t.root.payload
}

// Height returns the height (max depth) of the tree. Returns -1 if the tree
// has no nodes. A (rooted) tree with only a node (the root) has a height of
// zero.
func (t *Tree) Height() int {
	return int(height(t.root))
}

// Upsert inserts or updates an item.
func (t *Tree) Upsert(key int64, payload interface{}) {
	t.Lock()
	defer t.Unlock()

	if existing := search(t.root, key); existing != nil {
		existing.payload = payload

		return
	}

	var (
		parent  *node
		x       = t.root
		newNode = &node{
			key:     key,
			payload: payload,
		}
	)

	for x != nil {
		parent = x

		if newNode.key < parent.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	newNode.parent = parent

	switch {
	case parent == nil:
		t.root = newNode
	case newNode.key < parent.key:
		parent.left = newNode
	default:
		parent.right = newNode
	}
}

// Search searches for a for a node based on its key and returns the payload.
func (t *Tree) Search(key int64) interface{} {
	t.RLock()
	defer t.RUnlock()

	if t.root == nil {
		return nil
	}

	n := search(t.root, key)

	if n == nil {
		return nil
	}

	return n.payload
}

// Min returns the payload of the Node with the lowest key, or nil.
func (t *Tree) Min() interface{} {
	t.RLock()
	defer t.RUnlock()

	n := min(t.root)

	if n == nil {
		return nil
	}

	return n.payload
}

// Max returns the payload of the Node with the highest key, or nil.
func (t *Tree) Max() interface{} {
	t.RLock()
	defer t.RUnlock()

	n := max(t.root)

	if n == nil {
		return nil
	}

	return n.payload
}

// Successor returns the next highest neighbour (key-wise) of the Node with the
// passed key.
func (t *Tree) Successor(key int64) interface{} {
	t.RLock()
	defer t.RUnlock()

	n := successor(search(t.root, key))

	if n == nil {
		return nil
	}

	return n.payload
}

func height(node *node) float64 {
	if node == nil {
		return -1
	}

	return 1 + math.Max(height(node.left), height(node.right))
}

func successor(node *node) *node {
	if node.right != nil {
		return min(node.right)
	}

	parent := node.parent

	for parent != nil && node == parent.right {
		node = parent
		parent = node.parent
	}

	return parent
}

func max(node *node) *node {
	for node != nil && node.right != nil {
		node = node.right
	}

	return node
}

func min(node *node) *node {
	for node != nil && node.left != nil {
		node = node.left
	}

	return node
}

func search(node *node, key int64) *node {
	if node == nil || node.key == key {
		return node
	}

	for node != nil && node.key != key {
		if key > node.key {
			node = node.right
		} else {
			node = node.left
		}
	}

	return node
}
