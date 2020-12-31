// Package bst implements a binary search tree with arbitrary payloads.
package bst

import (
	"math"
	"sync"
)

// NewBSTree returns an empty binary search tree.
func NewBSTree() *BSTree {
	return &BSTree{
		RWMutex: sync.RWMutex{},
		root:    nil,
	}
}

// Root returns the payload of the root node of the tree.
func (t *BSTree) Root() interface{} {
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
func (t *BSTree) Height() int {
	t.RLock()
	defer t.RUnlock()

	return int(height(t.root))
}

// Upsert inserts or updates an item. Runs in O(lg n) time on average.
func (t *BSTree) Upsert(key int64, payload interface{}) {
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

// Search searches for a node based on its key and returns the payload.
func (t *BSTree) Search(key int64) interface{} {
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
func (t *BSTree) Min() interface{} {
	t.RLock()
	defer t.RUnlock()

	n := min(t.root)

	if n == nil {
		return nil
	}

	return n.payload
}

// Max returns the payload of the Node with the highest key, or nil.
func (t *BSTree) Max() interface{} {
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
func (t *BSTree) Successor(key int64) interface{} {
	t.RLock()
	defer t.RUnlock()

	n := successor(search(t.root, key))

	if n == nil {
		return nil
	}

	return n.payload
}

// Delete deletes a node with a given key. This runs in O(h) time with h being
// the height of the tree.
func (t *BSTree) Delete(key int64) {
	t.Lock()
	defer t.Unlock()

	if n := search(t.root, key); n != nil {
		t.delete(n)
	}
}

func (t *BSTree) delete(node *node) {
	switch {
	// If the node has no left subtree, replace it with its right subtree.
	case node.left == nil:
		t.transplant(node, node.right)

	// If the node has a left subtree but not a right one, replace it with
	// 	its right subtree.
	case node.right == nil:
		t.transplant(node, node.left)

	// Node has two children.
	default:
		// The node's successor must be the smallest key in the right subtree,
		// which has no left child.
		succ := min(node.right)

		// If the successor is the node's right child, the successor doesn't
		// have a left subtree. We replace the node with its right child (the
		// successor) and leave the latter's right subtree in tact.
		if succ.parent != node {
			t.transplant(succ, succ.right)

			succ.right = node.right
			succ.right.parent = succ
		}

		// If the successor is the node's right child, replace the parent of
		// the node by its successor, attaching node's left child.
		t.transplant(node, succ)
		succ.left = node.left
		succ.left.parent = node
	}
}

func height(node *node) float64 {
	if node == nil {
		return -1
	}

	return 1 + math.Max(height(node.left), height(node.right))
}

func successor(node *node) *node {
	if node == nil {
		return nil
	}

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

// transplant replaces one subtree of a node as a child of its parent, with
// another subtree.
func (t *BSTree) transplant(nodeA, nodeB *node) {
	if nodeA == nil {
		return
	}

	switch {
	// If nodeA is the root, nodeB will be root now.
	case nodeA.parent == nil:
		t.root = nodeB

	// If nodeA is a left-child, replace with nodeB.
	case nodeA == nodeA.parent.left:
		nodeA.parent.left = nodeB

	// If nodeA is a right-child, replace with nodeB.
	default:
		nodeA.parent.right = nodeB
	}

	// Update parent relationship.
	if nodeB != nil {
		nodeB.parent = nodeA.parent
	}
}
